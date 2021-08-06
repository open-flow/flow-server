package search

import "autoflow/pkg/entities/graph"

type FindActiveRequest struct {
	graph.DataEvent
}

type FindActiveResponse struct {
	Graphs []*ActiveGraph
}

type ActiveGraph struct {
	Graph  *graph.DBGraph
	Active []*graph.DBEventCard
}
