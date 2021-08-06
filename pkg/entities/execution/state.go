package execution

import (
	"autoflow/pkg/entities/graph"
)

type State struct {
	Graph      *graph.DBGraph     `json:"graph"`
	Card       *graph.DBEventCard `json:"card"`
	RawRequest interface{}        `json:"rawRequest"`

	Cursor *Cursor `json:"cursor"`
	Memory *Memory `json:"memory"`
}

type Cursor struct {
	Node *graph.DataNode `json:"node"`

	Current graph.DataConnection   `json:"current"`
	Path    []graph.DataConnection `json:"path"`
	Next    []graph.DataConnection `json:"next"`
}

type Memory struct {
	Context  map[string]interface{} `json:"context"`
	Response interface{}            `json:"response"`
}
