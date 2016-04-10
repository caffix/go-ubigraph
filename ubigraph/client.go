package ubigraph

import (
	"bytes"
	"errors"
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

func NewClient() (*Ubigraph, error) {
	rpcurl := "http://127.0.0.1:20738/RPC2"

	return &Ubigraph{rpcUrl: rpcurl}, nil
}

func (ubi *Ubigraph) Clear() error {
	method := "ubigraph.clear"

	status, err := ubi.Call(method, nil)
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.clear failed")
	}
	return nil
}

func (ubi *Ubigraph) SetURL(url string) error {
	ubi.rpcUrl = url
	return nil
}

func (ubi *Ubigraph) Call(method string, args interface{}) (status int, err error) {
	reply, err := xmlRpcCall(ubi.rpcUrl, method, args)
	if err != nil {
		return
	}
	status = reply.Status
	return
}

func xmlRpcCall(url string, method string, args interface{}) (reply result, err error) {
	buf, _ := xml.EncodeClientRequest(method, args)

	resp, err := http.Post(url, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	return
}
