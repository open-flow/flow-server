package state

import (
	"autoflow/pkg/storage/graph"
)

const LOOPING_MAX_COUNTER = 50

// swagger:model endpointState
type State struct {
	Graph      *graph.DBGraph     `json:"graph"`
	Card       *graph.DBEventCard `json:"card"`
	RawRequest interface{}        `json:"rawRequest"`

	Cursor *Cursor `json:"cursor"`
	Memory *Memory `json:"memory"`
}

func (s *State) GetProjectId() uint {
	return s.Graph.ProjectId
}
