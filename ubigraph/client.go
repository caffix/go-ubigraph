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

var (
	session *client
	once    sync.Once
)

type client struct {
	rpcUrl string
	mu     sync.RWMutex
}

type result struct {
	Status int
}

func ubigraphError(method string, status int) error {
	return fmt.Errorf("%s failed with status: %d", method, status)
}

// Client provides a reference to the singleton session object with a Ubigraph server.
// The function assumes the Ubigraph server is on the localhost.
// It returns the object for making API calls that manipulate the graph.
func Client() *client {
	once.Do(func() {
		session = &client{
			rpcUrl: "http://127.0.0.1:20738/RPC2",
		}
	})

	return session
}

// Clear removes all elements from the Ubigraph server.
func (c *client) Clear() error {
	method := "ubigraph.clear"

	status, err := c.call(method, nil)
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetURL changes the URL used to reach the XMLRPC Ubigraph server.
func (c *client) SetURL(url string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.rpcUrl = url
	return nil
}

// call will make a XMLRPC method call to the Ubigraph server.
// It returns the status integer resulting from the XMLRPC method call.
func (c *client) call(method string, args interface{}) (status int, err error) {
	var reply result

	buf, _ := xml.EncodeClientRequest(method, args)

	c.mu.Lock()
	url := c.rpcUrl
	c.mu.Unlock()
	resp, err := http.Post(url, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	status = reply.Status
	return
}
