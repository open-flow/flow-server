package batch

import (
	"autoflow/pkg/common"
	"autoflow/pkg/storage/graph"
)

type SaveRequest struct {
	common.ProjectModel

	Nodes       []graph.DBNode       `json:"nodes"`
	Cards       []graph.DBEventCard  `json:"cards"`
	Connections []graph.DBConnection `json:"connections"`
}

type SaveResponse struct {
	common.ProjectModel

	Nodes       []*graph.DBNode       `json:"nodes"`
	Cards       []*graph.DBEventCard  `json:"cards"`
	Connections []*graph.DBConnection `json:"connections"`
}

type DeleteRequest struct {
	common.ProjectModel

	Nodes       []uint64 `json:"nodes"`
	Cards       []uint64 `json:"cards"`
	Connections []uint64 `json:"connections"`
}
