// A Ubigraph client API and callback server.
// https://github.com/caffix/go-ubigraph/
// Licensed under the MIT license
package ubigraph

import (
	"bytes"
	"fmt"
	"github.com/caffix/gorilla-xmlrpc/xml"
	"net/http"
)

type Graph struct {
	rpcUrl          string
	callbackStarted bool
	vertexCallbacks map[VertexID]func(VertexID)
	styleCallbacks  map[VertexStyleID]func(VertexID)
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

var (
	graphs     map[string]*Graph
	serverMsgs chan serverMessage
)

func init() {
	// Graph objects are created for each host running a Ubigraph server
	graphs = make(map[string]*Graph)
	// start the message manager and channel
	serverMsgs = make(chan serverMessage)
	go msgManager()
}

func ubigraphError(method string, status int) error {
	return fmt.Errorf("%s failed with status: %d", method, status)
}

// Accepts ubigraph request messages over a channel and sends them
// out to the target ubigraph server.
// Runs as a separate goroutine.
func msgManager() {
	type result struct{ Status int }

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

func newGraph(addr string) *Graph {
	g := &Graph{rpcUrl: "http://" + addr + ":20738/RPC2"}

	g.vertexCallbacks = make(map[VertexID]func(VertexID))
	g.styleCallbacks = make(map[VertexStyleID]func(VertexID))
	graphs[addr] = g
	return g
}

// Ubigraph provides a reference to a session object with the Ubigraph server.
// The function assumes the Ubigraph server is on the localhost, unless the
// optional argument indicates otherwise.
// It returns the object for making API calls that manipulate the graph.
func Ubigraph(url ...string) *Graph {
	addr := "127.0.0.1"

	if l := len(url); l > 1 {
		return nil
	} else if l == 1 {
		a := graphAddr(url[0])

		if a == "" {
			return nil
		}
		addr = a
	}

	if g, ok := graphs[addr]; ok {
		return g
	}
	return newGraph(addr)
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

// call will make a XMLRPC method call to the Ubigraph server.
// It returns the status integer resulting from the XMLRPC method call.
func (g *Graph) serverCall(method string, args interface{}) (int, error) {
	var msg serverMessage
	var resp serverResponse

	msg.url = g.rpcUrl
	msg.buf, _ = xml.EncodeClientRequest(method, args)
	msg.resp = make(chan serverResponse)

	serverMsgs <- msg
	resp = <-msg.resp

	return resp.status, resp.err
}
