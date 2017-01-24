package main

import (
	"github.com/caffix/go-ubigraph/ubigraph"
	"time"
)

var N int = 10

func idx(i, j, k int) ubigraph.VertexID {
	return ubigraph.VertexID(i*N*N + j*N + k)
}

func main() {
	graph := ubigraph.Ubigraph()
	graph.Clear()

	sid, err := graph.NewVertexStyle(0)
	if err != nil {
		return
	}
	graph.SetVertexStyleAttribute(sid, "shape", "octahedron")
	graph.SetVertexStyleAttribute(sid, "size", "2.0")
	graph.SetVertexStyleAttribute(sid, "color", "#00ff00")

	for i := 0; i < N; i += 1 {
		for j := 0; j < N; j += 1 {
			for k := 0; k < N; k += 1 {
				v := idx(i, j, k)

				graph.NewVertexWithID(v)
				graph.ChangeVertexStyle(v, sid)
				if i != 0 {
					graph.NewEdge(idx(i-1, j, k), v)
				}
				if j != 0 {
					graph.NewEdge(idx(i, j-1, k), v)
				}
				if k != 0 {
					graph.NewEdge(idx(i, j, k-1), v)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
