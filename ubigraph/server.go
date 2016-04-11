package ubigraph

import (
	"github.com/caffix/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"net/http"
	"strconv"
	"strings"
)

func (ubi *Ubigraph) Callback(r *http.Request, args *struct{ VertexID int }, reply *struct{ Status int }) error {
	ubi.cbRoutine(args.VertexID)
	reply.Status = 0
	return nil
}

// SetCallbackServerAddr assigns the IP address that will be provided to the Ubigraph server for callback.
func (ubi *Ubigraph) SetCallbackServerAddr(ip string) {
	ubi.cbServerAddr = ip
}

// SetCallbackServerPort assigns the port number that will be provided to the Ubigraph server for callback.
func (ubi *Ubigraph) SetCallbackServerPort(port int) {
	ubi.cbServerPort = strconv.Itoa(port)
}

// SetCallbackRoutine assigns the Go function that will be executed as the vertex double-click callback.
func (ubi *Ubigraph) SetCallbackRoutine(fn func(int)) {
	ubi.cbRoutine = fn
}

// StartCallbackServer creates a XMLRPC over HTTP server listening for the Ubigraph callback.
// This method does not return on success, since it continues listening.
func (ubi *Ubigraph) StartCallbackServer() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	if err := RPC.RegisterService(ubi, ""); err == nil {
		http.Handle("/vertex_callback", RPC)
		http.ListenAndServe(":"+ubi.cbServerPort, nil)
	}
}

// SetVertexStyleCallback sets the double-click callback attribute for the identified style.
func (ubi *Ubigraph) SetVertexStyleCallback(styleID int) error {
	s := []string{"http://", ubi.cbServerAddr, ":", ubi.cbServerPort, "/vertex_callback/Ubigraph.Callback"}
	return ubi.SetVertexStyleAttribute(styleID, "callback_left_doubleclick", strings.Join(s, ""))
}
