package batch

import (
	"autoflow/pkg/entities/graph"
)

type SaveRequest struct {
	graph.IDGraph

	Nodes       []graph.DataNode
	Cards       []graph.DataEventCard
	Connections []graph.DataConnection
}

type SaveResponse struct {
	graph.IDGraph

	Nodes       []*graph.DBNode
	Cards       []*graph.DBEventCard
	Connections []*graph.DBConnection
}

type DeleteRequest struct {
	graph.IDGraph

	Nodes       []uint64
	Cards       []uint64
	Connections []uint64
}

type DeleteResponse struct {
}
