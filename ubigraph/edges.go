package ubigraph

type EdgeID int
type EdgeStyleID int

// NewEdge creates a edge on the graph connected to two vertices identified by arguments.
// It returns an Ubigraph server selected edge ID on success.
func (g *Graph) NewEdge(x, y VertexID) (EdgeID, error) {
	method := "ubigraph.new_edge"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2 int }{int(x), int(y)})
	if err != nil {
		return 0, err
	}
	return EdgeID(status), nil
}

// RemoveEdge deletes the edge with the identifier matching the argument.
func (g *Graph) RemoveEdge(id EdgeID) error {
	method := "ubigraph.remove_edge"

	status, err := g.serverCall(method, &struct{ Arg1 int }{int(id)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewEdgeWithID creates a edge on the graph connected to two selected vertices and with a chosen identifier.
func (g *Graph) NewEdgeWithID(id EdgeID, x, y VertexID) error {
	method := "ubigraph.new_edge_w_id"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2, Arg3 int }{int(id), int(x), int(y)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewEdgeStyle creates a edge style based on an existing style.
// It returns an Ubigraph server selected style ID on success.
func (g *Graph) NewEdgeStyle(parentStyle EdgeStyleID) (EdgeStyleID, error) {
	method := "ubigraph.new_edge_style"

	status, err := g.serverCall(method, &struct{ Arg1 int }{int(parentStyle)})
	if err != nil {
		return 0, err
	}
	return EdgeStyleID(status), nil
}

// NewEdgeStyleWithID creates a edge style with a chosen identifier based on an existing style.
func (g *Graph) NewEdgeStyleWithID(id, parentStyle EdgeStyleID) error {
	method := "ubigraph.new_edge_style_w_id"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2 int }{int(id), int(parentStyle)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// ChangeEdgeStyle changes the identified edge's style.
func (g *Graph) ChangeEdgeStyle(edge EdgeID, style EdgeStyleID) error {
	method := "ubigraph.change_edge_style"

	status, err := g.serverCall(method, &struct{ Arg1, Arg2 int }{int(edge), int(style)})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetEdgeAttribute modifies the attributes of the identified edge.
func (g *Graph) SetEdgeAttribute(id EdgeID, attribute, value string) error {
	method := "ubigraph.set_edge_attribute"

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

// SetEdgeStyleAttribute modifies the attributes of the identified edge style.
func (g *Graph) SetEdgeStyleAttribute(id EdgeStyleID, attribute, value string) error {
	method := "ubigraph.set_edge_style_attribute"

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
