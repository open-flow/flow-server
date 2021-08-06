package storage

import "autoflow/pkg/entities/graph"

type DeleteRequest struct {
	ProjectID uint
	ID        uint
}

type DeleteResponse struct {
}

type GetGraphRequest struct {
	ProjectID uint
	ID        uint
}

type GetGraphResponse struct {
	graph.DBGraph
}

type ListGraphRequest struct {
	ProjectIDs []uint64
}

type ListGraphResponse struct {
	Graphs []*graph.DBGraph
}
