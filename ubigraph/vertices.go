package ubigraph

// NewVertex creates a vertex on the graph.
// It returns an Ubigraph server selected vertex ID on success.
func (c *client) NewVertex() (int, error) {
	method := "ubigraph.new_vertex"

	return c.call(method, nil)
}

// RemoveVertex deletes the vertex with the identifier matching the argument.
func (c *client) RemoveVertex(vertID int) error {
	method := "ubigraph.remove_vertex"

	status, err := c.call(method, &struct{ Arg1 int }{vertID})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// NewVertexWithID creates a vertex on the graph with a chosen identifier.
func (c *client) NewVertexWithID(vertID int) error {
	method := "ubigraph.new_vertex_w_id"

	status, err := c.call(method, &struct{ Arg1 int }{vertID})
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
func (c *client) NewVertexStyle(parentStyle int) (int, error) {
	method := "ubigraph.new_vertex_style"

	status, err := c.call(method, &struct{ Arg1 int }{parentStyle})
	if err != nil {
		return 0, err
	}
	return status, nil
}

// NewVertexStyleWithID creates a vertex style with a chosen identifier based on an existing style.
func (c *client) NewVertexStyleWithID(styleID, parentStyle int) error {
	method := "ubigraph.new_vertex_style_w_id"

	status, err := c.call(method, &struct{ Arg1, Arg2 int }{styleID, parentStyle})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// ChangeVertexStyle changes the identified vertex's style.
func (c *client) ChangeVertexStyle(vertID, styleID int) error {
	method := "ubigraph.change_vertex_style"

	status, err := c.call(method, &struct{ Arg1, Arg2 int }{vertID, styleID})
	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetVertexAttribute modifies the attributes of the identified vertex.
func (c *client) SetVertexAttribute(vertID int, attribute, value string) error {
	method := "ubigraph.set_vertex_attribute"

	status, err := c.call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{vertID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return ubigraphError(method, status)
	}
	return nil
}

// SetVertexStyleAttribute modifies the attributes of the identified vertex style.
func (c *client) SetVertexStyleAttribute(styleID int, attribute, value string) error {
	method := "ubigraph.set_vertex_style_attribute"

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
