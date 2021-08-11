package state

import "autoflow/pkg/storage/graph"

type Cursor struct {
	Node *graph.DataNode `json:"node"`

	Current graph.DataConnection   `json:"current"`
	Path    []graph.DataConnection `json:"path"`
	Next    []graph.DataConnection `json:"next"`
}

func (c *Cursor) Module() string {
	return c.Node.Module
}

func (c *Cursor) Function() string {
	return c.Node.Function
}
