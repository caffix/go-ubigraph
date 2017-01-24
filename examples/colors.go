package main

import (
	"fmt"
	"github.com/caffix/go-ubigraph/ubigraph"
)

func main() {
	var i, N ubigraph.VertexID
	graph := ubigraph.Ubigraph()
	graph.Clear()

	N = 20
	for i = 0; i < N; i += 1 {
		graph.NewVertexWithID(i)
	}

	for i = 0; i < N; i += 1 {
		var r int = int(float32(i) / float32(N) * 255)
		c := fmt.Sprintf("#%02x%02x%02x", r, 255-r, 255)
		graph.SetVertexAttribute(i, "color", c)
		graph.NewEdge(i, (i+1)%(N/2))
	}
}
