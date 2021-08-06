package execution

import (
	"autoflow/pkg/entities/graph"
)

type State struct {
	Graph      *graph.DBGraph
	Card       *graph.DBEventCard
	RawRequest interface{}

	Cursor *Cursor
	Memory *Memory
}

type Cursor struct {
	Node *graph.DataNode

	Current graph.DataConnection
	Path    []graph.DataConnection
	Next    []graph.DataConnection
}

type Memory struct {
	Context  map[string]interface{}
	Response interface{}
}
