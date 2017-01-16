// A Ubigraph client API and callback server.
// https://github.com/caffix/go-ubigraph/
// Licensed under the MIT license
package ubigraph

import (
	"bytes"
	"fmt"
	"github.com/caffix/gorilla-xmlrpc/xml"
	"net/http"
	"sync"
)

var session = &Graph{
	rpcUrl: "http://127.0.0.1:20738/RPC2",
}

var serverMsgs chan serverMessage

type Graph struct {
	rpcUrl string
	mu     sync.RWMutex
}

type serverMessage struct {
	url  string
	buf  []byte
	resp chan serverResponse
}

type serverResponse struct {
	err    error
	status int
}

type result struct {
	Status int
}

func init() {
	serverMsgs = make(chan serverMessage)

	go sendServerMessages()
}

func ubigraphError(method string, status int) error {
	return fmt.Errorf("%s failed with status: %d", method, status)
}

func sendServerMessages() {
	for {
		select {
		case msg := <-serverMsgs:
			var response serverResponse

			resp, err := http.Post(msg.url, "text/xml", bytes.NewBuffer(msg.buf))
			response.err = err
			if response.err == nil {
				var reply result

				response.err = xml.DecodeClientResponse(resp.Body, &reply)
				if response.err == nil {
					response.status = reply.Status
				}

				resp.Body.Close()
			}

			msg.resp <- response
		}
	}
}

// Ubigraph provides a reference to the singleton session object with a Ubigraph server.
// The function assumes the Ubigraph server is on the localhost.
// It returns the object for making API calls that manipulate the graph.
func Ubigraph() *Graph {
	return session
}

// Clear removes all elements from the Ubigraph server.
func (g *Graph) Clear() error {
	method := "ubigraph.clear"

	status, err := g.serverCall(method, nil)
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

func (g *Graph) GetURL() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.rpcUrl
}

// SetURL changes the URL used to reach the XMLRPC Ubigraph server.
func (g *Graph) SetURL(url string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.rpcUrl = url
}

// call will make a XMLRPC method call to the Ubigraph server.
// It returns the status integer resulting from the XMLRPC method call.
func (g *Graph) serverCall(method string, args interface{}) (int, error) {
	var msg serverMessage
	var resp serverResponse

	msg.url = g.GetURL()
	msg.buf, _ = xml.EncodeClientRequest(method, args)
	msg.resp = make(chan serverResponse)

	serverMsgs <- msg
	resp = <-msg.resp

	return resp.status, resp.err
}
