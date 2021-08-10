package engine

import (
	"autoflow/pkg/entities/graph"
)

type State struct {
	Graph      *graph.DBGraph     `json:"graph"`
	Card       *graph.DBEventCard `json:"card"`
	RawRequest interface{}        `json:"rawRequest"`

	*Cursor `json:"cursor"`
	*Memory `json:"memory"`
}

func (s *State) GetProjectId() uint {
	return s.Graph.ProjectId
}
