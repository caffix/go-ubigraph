package ubigraph

import (
	"github.com/divan/gorilla-xmlrpc/xml"
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

func (ubi *Ubigraph) SetCallbackServerAddr(ip string) {
	ubi.cbServerAddr = ip
}

func (ubi *Ubigraph) SetCallbackServerPort(port int) {
	ubi.cbServerPort = strconv.Itoa(port)
}

func (ubi *Ubigraph) SetCallbackRoutine(fn func(int)) {
	ubi.cbRoutine = fn
}

func (ubi *Ubigraph) StartCallbackServer() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	if err := RPC.RegisterService(ubi, ""); err == nil {
		http.Handle("/vertex_callback", RPC)
		http.ListenAndServe(":"+ubi.cbServerPort, nil)
	}
}

func (ubi *Ubigraph) SetVertexStyleCallback(styleID int) error {
	s := []string{"http://", ubi.cbServerAddr, ":", ubi.cbServerPort, "/vertex_callback/Ubigraph.Callback"}
	return ubi.SetVertexStyleAttribute(styleID, "callback_left_doubleclick", strings.Join(s, ""))
}
