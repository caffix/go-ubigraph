# go-ubigraph

## Overview

go-ubigraph is an implementation of the Ubigraph client API and callback server.

## Installation

The Ubigraph server must already be installed before using this API. Unfortunately, the original website (http://www.ubietylab.net/ubigraph) where the binaries could be downloaded is no longer available. Assuming you have acquired the server software anyway, continue to obtain the Go API package as follows:

    go get "github.com/caffix/go-ubigraph"

That's it! You're ready to roll :-)

## Usage

**Simple example to create two vertices and connect them with an edge.**

    import "github.com/caffix/go-ubigraph/ubigraph"

    graph := ubigraph.Client()
    graph.Clear()
    
    x, _ := graph.NewVertex()
    y, _ := graph.NewVertex()
    graph.NewEdge(x, y)


**Changing attributes of a vertex.**

    if id, err := graph.NewVertex(); err == nil {
        graph.SetVertexAttribute(id, "color", "#ff0000")
        graph.SetVertexAttribute(id, "shape", "sphere")
    }


**Creating a vertex style and modifying its attributes.**

    //The new style is based on the default style identified by '0'
    sid, _ := graph.NewVertexStyle(0)

    //Modify the vertex attributes within the style identified by 'sid'
    graph.SetVertexStyleAttribute(sid, "shape", "torus")
    graph.SetVertexStyleAttribute(sid, "size", "2.0")
    graph.SetVertexStyleAttribute(sid, "color", "#00ffff")

    //Create a new vertex and change its style
    if id, err := graph.NewVertex(); err == nil {
        graph.ChangeVertexStyle(id, sid)
    }

**Vertex Attributes of the Ubigraph Server.**

| Attribute | Values | Default |
| :------------- | :------------- | :-------------: |
| color | String of the form "#000000" specifying an rgb triple. | `"#0000ff"` |
| shape | cone, cube, dodecahedron, icosahedron, octahedron, sphere, and torus. | `"cube"` |
| shapedetail | Indicates the level of detail with which the shape should be rendered. This is relevant only for the sphere, cone, and torus shapes, which are described by polygons. Performance may improve for large graphs if the level of detail is reduced. Sensible values from 5 to 40. If shapedetail=0, the level of detail varies with the framerate. | `"10"` |
| label | A string to be displayed near the vertex. | `""` |
| size | Real number indicating the relative size of the shape. This is for rendering only, and does not affect layout. | `"1.0"` |
| fontcolor | String of the form "#000000" specifying an rgb triple. | `"#ffffff"` |
| fontfamily | String indicating the font to be used for the label. Recognized choices are "Helvetica" and "Times Roman". Only the combinations of family and size shown below are recognized; other choices of family and size result in a best guess. | `"Helvetica"` |
| fontsize | Integer giving the size of the font, in points, used for the label. | `"12"` |
| visible | Whether this vertex is drawn. | `"true"` |


**Edge Attributes of the Ubigraph Server.**

| Attribute | Values | Default |
| :------------- | :------------- | :-------------: |
| arrow | If true, an arrowhead is drawn. | `"false"` |
| arrow_position | On an edge (x,y), if arrow_position=1.0 then the arrowhead is drawn so that the tip is touching y. If arrow_position=0.0 the beginning of the arrowhead is touching x. If arrow_position=0.5 the arrowhead is midway between the two vertices. | `"0.5"` |
| arrow_radius | How thick the arrowhead is. | `"1.0"` |
| arrow_length | How long the arrowhead is. | `"1.0"` |
| arrow_reverse | If true, the arrowhead on an edge (x,y) will point toward x. | `"false"` |
| color, label, fontcolor, fontfamily, fontsize, visible | See vertex style attributes. |  |
| oriented | If true, the edge tries to point 'downward'. | `"false"` |
| spline | If true, a curved edge is rendered. A curved edge tries to avoid other curved edges in the layout, which can result in cleaner-looking layouts. | `"false"` |
| showstrain | If true, edges are colored according to their relative length. Longer than average edges are drawn in red. Edges of average length are drawn in white. Shorter than average edges are drawn in blue. | `"false"` |
| stroke | The stroke style to be used: one of "solid", "dashed", "dotted", or "none". If the "none" style is used, no line is drawn. However, any decorations of the edge, e.g., arrowhead and label, will be drawn. | `"solid"` |
| strength | How much the edge will pull its vertices together. For edges that are drawn but do not affect layout, use "0.0". | `"1.0"` |
| width | How wide the edge is. | `"1.0"` |


## License

This package is licensed under the MIT license. See LICENSE for details.
