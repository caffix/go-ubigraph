package main

import (
	"fmt"
	"github.com/caffix/go-ubigraph/ubigraph"
	"time"
)

func callback(id int) {
	fmt.Printf("Left Double-Click for Vertex: %d\n", id)
}

func main() {
	graph := ubigraph.Client()
	graph.Clear()

	sytleID, err := graph.NewVertexStyle(0)
	if err != nil {
		return
	}
	graph.SetVertexStyleAttribute(sytleID, "shape", "sphere")
	graph.SetVertexStyleAttribute(sytleID, "color", "#ff0000")

	go func() {
		cirlen := 10
		vertices := make([]int, cirlen, cirlen)
		for i := 0; i < cirlen; i = i + 1 {
			vertices[i], _ = graph.NewVertex()
			graph.ChangeVertexStyle(vertices[i], sytleID)
		}

		for i := 0; i < cirlen; i = i + 1 {
			time.Sleep(time.Second)
			if i != 0 {
				graph.NewEdge(vertices[i-1], vertices[i])
			}
			graph.NewEdge(vertices[i], vertices[(i+1)%cirlen])
		}
	}()
	server := ubigraph.CallbackServer()
	server.SetCallbackRoutine(callback)
	graph.SetVertexStyleCallback(sytleID, server)
	server.Start()
}
