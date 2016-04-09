package ubigraph

import (
	"errors"
)

func (ubi *Ubigraph) NewEdge(vertIDX int, vertIDY int) (int, error) {
	method := "ubigraph.new_edge"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{vertIDX, vertIDY})
	if err != nil {
		return 0, err
	}
	return status, nil
}

func (ubi *Ubigraph) RemoveEdge(edgeID int) error {
	method := "ubigraph.remove_edge"

	status, err := ubi.Call(method, &struct{ Arg1 int }{edgeID})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.remove_edge failed")
	}
	return nil
}

func (ubi *Ubigraph) NewEdgeWithID(vertIDX int, vertIDY int) error {
	method := "ubigraph.new_edge_w_id"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{vertIDX, vertIDY})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.new_edge_w_id failed")
	}
	return nil
}

func (ubi *Ubigraph) NewEdgeStyle(parentStyle int) (int, error) {
	method := "ubigraph.new_edge_style"

	status, err := ubi.Call(method, &struct{ Arg1 int }{parentStyle})
	if err != nil {
		return 0, err
	}
	return status, nil
}

func (ubi *Ubigraph) NewEdgeStyleWithID(styleID int, parentStyle int) error {
	method := "ubigraph.new_edge_style_w_id"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{styleID, parentStyle})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.new_edge_style_w_id failed")
	}
	return nil
}

func (ubi *Ubigraph) ChangeEdgeStyle(edgeID int, styleID int) error {
	method := "ubigraph.change_edge_style"

	status, err := ubi.Call(method, &struct{ Arg1, Arg2 int }{edgeID, styleID})
	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.change_edge_style failed")
	}
	return nil
}

func (ubi *Ubigraph) SetEdgeAttribute(edgeID int, attribute string, value string) error {
	method := "ubigraph.set_edge_attribute"

	status, err := ubi.Call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{edgeID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.set_edge_attribute failed")
	}
	return nil
}

func (ubi *Ubigraph) SetEdgeStyleAttribute(styleID int, attribute string, value string) error {
	method := "ubigraph.set_edge_style_attribute"

	status, err := ubi.Call(method,
		&struct {
			Arg1       int
			Arg2, Arg3 string
		}{styleID, attribute, value})

	if err != nil {
		return err
	}
	if status != 0 {
		return errors.New("ubigraph.set_edge_style_attribute failed")
	}
	return nil
}
