package batch

import (
	"autoflow/pkg/entities/common"
	"autoflow/pkg/entities/graph"
)

type SaveRequest struct {
	common.IDProject

	Nodes       []graph.DBNode       `json:"nodes"`
	Cards       []graph.DBEventCard  `json:"cards"`
	Connections []graph.DBConnection `json:"connections"`
}

type SaveResponse struct {
	common.IDProject

	Nodes       []*graph.DBNode       `json:"nodes"`
	Cards       []*graph.DBEventCard  `json:"cards"`
	Connections []*graph.DBConnection `json:"connections"`
}

type DeleteRequest struct {
	common.IDProject

	Nodes       []uint64 `json:"nodes"`
	Cards       []uint64 `json:"cards"`
	Connections []uint64 `json:"connections"`
}

type DeleteResponse struct {
}
