package dtos

import "autoflow/pkg/orm"

type ActiveEvent struct {
	Platform string

	OwnerType string
	OwnerId   string

	ResourceType string
	ResourceId   string

	ContextType string
	ContextId   string

	InitiatorType string
	InitiatorId   string

	StaticType string
	StaticId   string
}

type FindActiveGraphResponse struct {
	Graphs []*ActiveGraph
}

type ActiveGraph struct {
	Graph       *orm.Graph
	ActiveCards []*orm.EventCard
}
