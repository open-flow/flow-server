package batch

import (
	"autoflow/pkg/entities/graph"
)

type SaveRequest struct {
	ProjectID uint
	ID        uint

	Nodes       []graph.DataNode
	Cards       []graph.DataEventCard
	Connections []graph.DataConnection
}

type SaveResponse struct {
	ProjectID uint
	ID        uint

	Nodes       []*graph.DBNode
	Cards       []*graph.DBEventCard
	Connections []*graph.DBConnection
}

type DeleteRequest struct {
	ProjectID uint
	ID        uint

	Nodes       []uint64
	Cards       []uint64
	Connections []uint64
}

type DeleteResponse struct {
}
