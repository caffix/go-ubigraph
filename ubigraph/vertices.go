package ubigraph

type VertexID int
type VertexStyleID int

// NewVertex creates a vertex on the graph.
// It returns an Ubigraph server selected vertex ID on success.
func (g *Graph) NewVertex() (VertexID, error) {
	method := "ubigraph.new_vertex"

	status, err := g.serverCall(method, nil)
	if err != nil {
		return 0, err
	}
	return VertexID(status), nil
}

// RemoveVertex deletes the vertex with the identifier matching the argument.
func (g *Graph) RemoveVertex(id VertexID) error {
	method := "ubigraph.remove_vertex"

	status, err := g.serverCall(method, &struct{ Arg1 int }{int(id)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewVertexWithID creates a vertex on the graph with a chosen identifier.
func (g *Graph) NewVertexWithID(id VertexID) error {
	method := "ubigraph.new_vertex_w_id"

	status, err := g.serverCall(method, &struct{ Arg1 int }{int(id)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewVertexStyle creates a vertex style based on an existing style.
// It returns an Ubigraph server selected style ID on success.
func (g *Graph) NewVertexStyle(parentStyle VertexStyleID) (VertexStyleID, error) {
	method := "ubigraph.new_vertex_style"

	status, err := g.serverCall(method, &struct{ Arg1 int }{int(parentStyle)})
	if err != nil {
		return 0, err
	}
	return VertexStyleID(status), nil
}

// NewVertexStyleWithID creates a vertex style with a chosen identifier based on an existing style.
func (g *Graph) NewVertexStyleWithID(id, parentStyle VertexStyleID) error {
	method := "ubigraph.new_vertex_style_w_id"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2 int }{int(id), int(parentStyle)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// ChangeVertexStyle changes the identified vertex's style.
func (g *Graph) ChangeVertexStyle(vertex VertexID, style VertexStyleID) error {
	method := "ubigraph.change_vertex_style"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2 int }{int(vertex), int(style)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetVertexAttribute modifies the attributes of the identified vertex.
func (g *Graph) SetVertexAttribute(id VertexID, attribute, value string) error {
	method := "ubigraph.set_vertex_attribute"

	status, err := g.serverCall(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{int(id), attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetVertexStyleAttribute modifies the attributes of the identified vertex style.
func (g *Graph) SetVertexStyleAttribute(id VertexStyleID, attribute, value string) error {
	method := "ubigraph.set_vertex_style_attribute"

	status, err := g.serverCall(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{int(id), attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}
