package ubigraph

// NewEdge creates a edge on the graph connected to two vertices identified by arguments.
// It returns an Ubigraph server selected edge ID on success.
func (c *client) NewEdge(vertIDX, vertIDY int) (int, error) {
	method := "ubigraph.new_edge"

	status, err := c.call(method, &struct{ Arg1, Arg2 int }{vertIDX, vertIDY})
	if err != nil {
		return 0, err
	}
	return status, nil
}

// RemoveEdge deletes the edge with the identifier matching the argument.
func (c *client) RemoveEdge(edgeID int) error {
	method := "ubigraph.remove_edge"

	status, err := c.call(method, &struct{ Arg1 int }{edgeID})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewEdgeWithID creates a edge on the graph connected to two selected vertices and with a chosen identifier.
func (c *client) NewEdgeWithID(edgeID, vertIDX, vertIDY int) error {
	method := "ubigraph.new_edge_w_id"

	status, err := c.call(method, &struct{ Arg1, Arg2, Arg3 int }{edgeID, vertIDX, vertIDY})
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
func (c *client) NewEdgeStyle(parentStyle int) (int, error) {
	method := "ubigraph.new_edge_style"

	status, err := c.call(method, &struct{ Arg1 int }{parentStyle})
	if err != nil {
		return 0, err
	}
	return status, nil
}

// NewEdgeStyleWithID creates a edge style with a chosen identifier based on an existing style.
func (c *client) NewEdgeStyleWithID(styleID, parentStyle int) error {
	method := "ubigraph.new_edge_style_w_id"

	status, err := c.call(method, &struct{ Arg1, Arg2 int }{styleID, parentStyle})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// ChangeEdgeStyle changes the identified edge's style.
func (c *client) ChangeEdgeStyle(edgeID, styleID int) error {
	method := "ubigraph.change_edge_style"

	status, err := c.call(method, &struct{ Arg1, Arg2 int }{edgeID, styleID})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetEdgeAttribute modifies the attributes of the identified edge.
func (c *client) SetEdgeAttribute(edgeID int, attribute, value string) error {
	method := "ubigraph.set_edge_attribute"

	status, err := c.call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{edgeID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetEdgeStyleAttribute modifies the attributes of the identified edge style.
func (c *client) SetEdgeStyleAttribute(styleID int, attribute, value string) error {
	method := "ubigraph.set_edge_style_attribute"

	status, err := c.call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{styleID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}
