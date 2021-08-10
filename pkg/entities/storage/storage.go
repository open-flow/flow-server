package storage

import (
	"autoflow/pkg/entities/graph"
)

type ListGraphRequest struct {
	ProjectId []uint `json:"projectID" form:"projectId"`
}

type ListGraphResponse struct {
	Graphs []graph.DBGraph `json:"graphs"`
}
