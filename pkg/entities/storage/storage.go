package storage

import "autoflow/pkg/entities/graph"

type DeleteRequest struct {
	graph.IDProject
}

type DeleteResponse struct {
}

type GetGraphRequest struct {
	graph.IDProject
}

type GetGraphResponse struct {
	graph.DBGraph
}

type ListGraphRequest struct {
	ProjectId []uint `json:"projectID" form:"projectId"`
}

type ListGraphResponse struct {
	Graphs []graph.DBGraph `json:"graphs"`
}
