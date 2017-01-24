package ubigraph

import (
	"fmt"
	"github.com/caffix/gorilla-xmlrpc/xml"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Returns the IP address of the machine identified by the URL provided.
func graphAddr(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		return ""
	}
	// we are not interested in the port
	end := strings.Index(u.Host, ":")
	if end == -1 {
		end = len(u.Host)
	}
	// get the IP address
	addrs, err := net.LookupHost(u.Host[:end])
	if err != nil {
		return ""
	}
	return addrs[0]
}

// Returns the local IP address used as a source address when
// sending traffic to the address provided by the parameter.
func localAddr(addr string) (string, bool) {
	dst := net.UDPAddr{Port: 9} // port number selected is irrelevant
	dst.IP = net.ParseIP(addr)

	// setup a connection (no packets to be sent)
	c, err := net.DialUDP("udp", nil, &dst)
	if err != nil {
		return "", false
	}
	defer c.Close()
	// pull the source address
	src, ok := c.LocalAddr().(*net.UDPAddr)
	if !ok {
		return "", false
	}
	return src.IP.String(), true
}

// Creates a XMLRPC over HTTP server listening for the Ubigraph callback.
// This function does not return on success, since it continues listening.
func (g *Graph) startCallbackServer() {
	RPC := rpc.NewServer()

	// ubigraph uses XML
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")

	// register this graph object for handling the RPC calls
	if err := RPC.RegisterService(g, ""); err == nil {
		routes := mux.NewRouter()
		addr, _ := localAddr(graphAddr(g.rpcUrl))

		// setup the router for pulling out the style id from the URL
		routes.Handle("/style_callback/{id:[0-9]+}", RPC)
		routes.Handle("/vertex_callback", RPC)
		http.Handle("/", routes)
		http.ListenAndServe(addr+":20740", nil)
	}
}

func (g *Graph) checkCallbackServer() {
	if g.callbackStarted {
		return
	}
	// TODO: dynamically select the port number here
	g.callbackStarted = true
	go g.startCallbackServer()
}

// Called by the RPC library
func (g *Graph) Callback(r *http.Request, args *struct{ ID int }, reply *struct{ Status int }) error {
	vid := VertexID(args.ID)
	vars := mux.Vars(r)
	var cb func(VertexID)

	// look for a vertex callback routine
	if vars == nil {
		cb = g.vertexCallbacks[vid]
	} else if sid, ok := vars["id"]; ok {
		// a callback routine was registered with the vertex style
		i, err := strconv.Atoi(sid)
		if err != nil {
			reply.Status = 1
			return err
		}

		cb = g.styleCallbacks[VertexStyleID(i)]
	}
	// check if a user-registered callback routine was found
	if cb == nil {
		reply.Status = 1
		return fmt.Errorf("Callback for vertex %d was not found", int(vid))
	}
	// call the user-registered callback routine
	cb(vid)
	return nil
}

// SetVertexCallback sets the double-click callback attribute for the identified vertex
func (g *Graph) SetVertexCallback(id VertexID, f func(VertexID)) error {
	g.checkCallbackServer()
	g.vertexCallbacks[id] = f

	addr, _ := localAddr(graphAddr(g.rpcUrl))
	pieces := []string{"http://", addr, ":20740", "/vertex_callback/Graph.Callback"}
	url := strings.Join(pieces, "")

	return g.SetVertexAttribute(id, "callback_left_doubleclick", url)
}

// SetVertexStyleCallback sets the double-click callback attribute for the identified style
func (g *Graph) SetVertexStyleCallback(id VertexStyleID, f func(VertexID)) error {
	g.checkCallbackServer()
	g.styleCallbacks[id] = f

	addr, _ := localAddr(graphAddr(g.rpcUrl))
	sid := strconv.Itoa(int(id))
	pieces := []string{"http://", addr, ":20740", "/style_callback/" + sid + "/Graph.Callback"}
	url := strings.Join(pieces, "")

	return g.SetVertexStyleAttribute(id, "callback_left_doubleclick", url)
}
