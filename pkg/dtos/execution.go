package dtos

import (
	"autoflow/pkg/orm"
)

type ExecutionCursor struct {
	Visited    []Visited
	TargetId   uint64
	Graph      *orm.Graph
	EventCard  *orm.EventCard
	Context    map[string]interface{}
	RawRequest interface{}
}

type Visited struct {
	Connection orm.Connection
}
