package ubigraph

import (
	"errors"
)

func (ubi *Ubigraph) NewVertex() (int, error) {
	method := "ubigraph.new_vertex"

	return ubi.Call(method, nil)
}

func (ubi *Ubigraph) RemoveVertex(vertID int) error {
	method := "ubigraph.remove_vertex"

	status, err := ubi.Call(method, &struct{ Arg1 int }{vertID})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.remove_vertex failed")
	}
	return nil
}

func (ubi *Ubigraph) NewVertexWithID(vertID int) error {
	method := "ubigraph.new_vertex_w_id"

	status, err := ubi.Call(method, &struct{ Arg1 int }{vertID})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.new_vertex_w_id failed")
	}
	return nil
}

func (ubi *Ubigraph) NewVertexStyle(parentStyle int) (int, error) {
	method := "ubigraph.new_vertex_style"

	status, err := ubi.Call(method, &struct{ Arg1 int }{parentStyle})
	if err != nil {
		return 0, err
	}
	return status, nil
}

func (ubi *Ubigraph) ChangeVertexStyle(vertID int, styleID int) error {
	method := "ubigraph.change_vertex_style"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{vertID, styleID})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.change_vertex_style failed")
	}
	return nil
}

func (ubi *Ubigraph) NewVertexStyleWithID(styleID int, parentStyle int) error {
	method := "ubigraph.new_vertex_style_w_id"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{styleID, parentStyle})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.new_vertex_style_w_id failed")
	}
	return nil
}

func (ubi *Ubigraph) SetVertexAttribute(vertID int, attribute string, value string) error {
	method := "ubigraph.set_vertex_attribute"

	status, err := ubi.Call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{vertID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.set_vertex_attribute failed")
	}
	return nil
}

func (ubi *Ubigraph) SetVertexStyleAttribute(styleID int, attribute string, value string) error {
	method := "ubigraph.set_vertex_style_attribute"

	status, err := ubi.Call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{styleID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.set_vertex_style_attribute failed")
	}
	return nil
}
