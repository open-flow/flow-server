package search

import "autoflow/pkg/storage/graph"

type FindActiveRequest struct {
	graph.DataEvent
}

type FindActiveResponse struct {
	Graphs []*ActiveGraph `json:"graphs"`
}

type ActiveGraph struct {
	Graph  *graph.DBGraph       `json:"graph"`
	Active []*graph.DBEventCard `json:"active"`
}
