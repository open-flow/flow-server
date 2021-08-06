package storage

import "autoflow/pkg/entities/graph"

type DeleteRequest struct {
	graph.IDGraph
}

type DeleteResponse struct {
}

type GetGraphRequest struct {
	graph.IDGraph
}

type GetGraphResponse struct {
	graph.DBGraph
}

type ListGraphRequest struct {
	ProjectIDs []uint64 `json:"projectIDs"`
}

type ListGraphResponse struct {
	Graphs []*graph.DBGraph `json:"graphs"`
}
