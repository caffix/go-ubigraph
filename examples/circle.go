package main

import (
	"fmt"
	"github.com/caffix/go-ubigraph/ubigraph"
	"time"
)

func callback(id ubigraph.VertexID) {
	fmt.Printf("Left Double-Click for Vertex: %d\n", id)
}

func main() {
	wait := make(chan bool)
	graph := ubigraph.Ubigraph()
	graph.Clear()

	sid, err := graph.NewVertexStyle(0)
	if err != nil {
		return
	}
	graph.SetVertexStyleAttribute(sid, "shape", "sphere")
	graph.SetVertexStyleAttribute(sid, "color", "#ff0000")

	go func() {
		cirlen := 10
		vertices := make([]ubigraph.VertexID, cirlen, cirlen)

		// create all the vertices and assign them the new style
		for i := 0; i < cirlen; i = i + 1 {
			vertices[i], _ = graph.NewVertex()
			graph.ChangeVertexStyle(vertices[i], sid)
		}

		// connect one vertex to another each second until they form a circle
		for i := 0; i < cirlen; i = i + 1 {
			time.Sleep(time.Second)
			if i != 0 {
				graph.NewEdge(vertices[i-1], vertices[i])
			}
			graph.NewEdge(vertices[i], vertices[(i+1)%cirlen])
		}
	}()

	// assign a callback routine to all the vertices through the vertex style
	graph.SetVertexStyleCallback(sid, callback)
	// wait indefinitely for user activity with the ubigraph
	<-wait
}
