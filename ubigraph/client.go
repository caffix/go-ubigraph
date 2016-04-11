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

type Ubigraph struct {
	rpcUrl       string
	cbServerAddr string
	cbServerPort string
	cbRoutine    func(int)
}

type result struct {
	Status int
}

func ubigraphError(method string, status int) error {
	return fmt.Errorf("%s failed with status: %d", method, status)
}

// NewClient creates a unique session with a Ubigraph server.
// The function assumes the Ubigraph server is on the localhost.
// It returns the client object for making the API calls.
func NewClient() (*Ubigraph, error) {
	rpcurl := "http://127.0.0.1:20738/RPC2"

	return &Ubigraph{rpcUrl: rpcurl}, nil
}

// Clear removes all elements from the Ubigraph server.
func (ubi *Ubigraph) Clear() error {
	method := "ubigraph.clear"

	status, err := ubi.call(method, nil)
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetURL changes the URL used to reach the XMLRPC Ubigraph server.
func (ubi *Ubigraph) SetURL(url string) error {
	ubi.rpcUrl = url
	return nil
}

// Call will make a XMLRPC method call to the Ubigraph server.
// It returns the status integer resulting from the XMLRPC method call.
func (ubi *Ubigraph) call(method string, args interface{}) (status int, err error) {
	var reply result

	buf, _ := xml.EncodeClientRequest(method, args)

	resp, err := http.Post(ubi.rpcUrl, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	if err != nil {
		return
	}
	status = reply.Status
	return
}
