package storage

import "autoflow/pkg/orm"

type RequestFindActiveGraph struct {
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

type ActiveGraph struct {
	Graph       *orm.Graph
	ActiveCards []*orm.EventCard
}

type ResponseFindActiveGraph struct {
	Graphs []*ActiveGraph
}
