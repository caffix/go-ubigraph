package ubigraph

import (
	"github.com/caffix/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	server    *Callback
	oneServer sync.Once
)

type Callback struct {
	addr    string
	port    string
	routine func(int)
	mu      sync.RWMutex
}

// CallbackServer provides a reference to the singleton object that handles double-click callbacks from the Ubigraph server.
// The function assumes the Ubigraph and callback server are on the localhost.
// It returns the object for making callback server related API calls.
func CallbackServer() *Callback {
	oneServer.Do(func() {
		server = &Callback{
			addr: "127.0.0.1",
			port: "20740",
		}
	})

	return server
}

func (c *Callback) Proc(r *http.Request, args *struct{ VertexID int }, reply *struct{ Status int }) error {
	c.mu.Lock()
	cb := c.routine
	c.mu.Unlock()
	cb(args.VertexID)
	reply.Status = 0
	return nil
}

// SetCallbackServerAddr assigns the IP address that will be provided to the Ubigraph server for callback.
func (c *Callback) SetCallbackServerAddr(ip string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.addr = ip
}

// SetCallbackServerPort assigns the port number that will be provided to the Ubigraph server for callback.
func (c *Callback) SetCallbackServerPort(port int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.port = strconv.Itoa(port)
}

// SetCallbackRoutine assigns the Go function that will be executed as the vertex double-click callback.
func (c *Callback) SetCallbackRoutine(fn func(int)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.routine = fn
}

// Start creates a XMLRPC over HTTP server listening for the Ubigraph callback.
// This method does not return on success, since it continues listening.
func (c *Callback) Start() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	if err := RPC.RegisterService(c, ""); err == nil {
		http.Handle("/vertex_callback", RPC)
		c.mu.Lock()
		addr := c.addr
		port := c.port
		c.mu.Unlock()
		http.ListenAndServe(addr+":"+port, nil)
	}
}

// SetVertexCallback sets the double-click callback attribute for the identified vertex.
func (c *client) SetVertexCallback(vertID int, cb *Callback) error {
	cb.mu.Lock()
	pieces := []string{"http://", cb.addr, ":", cb.port, "/vertex_callback/Callback.Proc"}
	url := strings.Join(pieces, "")
	cb.mu.Unlock()
	return c.SetVertexAttribute(vertID, "callback_left_doubleclick", url)
}

// SetVertexStyleCallback sets the double-click callback attribute for the identified style.
func (c *client) SetVertexStyleCallback(styleID int, cb *Callback) error {
	cb.mu.Lock()
	pieces := []string{"http://", cb.addr, ":", cb.port, "/vertex_callback/Callback.Proc"}
	url := strings.Join(pieces, "")
	cb.mu.Unlock()
	return c.SetVertexStyleAttribute(styleID, "callback_left_doubleclick", url)
}
