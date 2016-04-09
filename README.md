# go-ubigraph

## Overview

go-ubigraph is an implementation of the Ubigraph client API and callback server.

## Installation

The Ubigraph server must already be installed before using this API. Unfortunately, the original website where the binaries could be downloaded is no longer available. Assuming you have acquired the server software anyway, continue to obtain the Go API package as follows:

    go get "github.com/caffix/go-ubigraph/ubigraph"

That's it! You're ready to roll :-)

## Usage

Simple example to create two vertices and connect them with an edge.

    graph, _ := ubigraph.NewClient()
    graph.Clear()
    
    x, _ := graph.NewVertex()
    y, _ := graph.NewVertex()
    graph.NewEdge(x, y)

See the circle example for a few additional details.

## Roadmap

1. Polish up the vertices and edges routines.

## License

This package is licensed under MIT license. See LICENSE for details.
