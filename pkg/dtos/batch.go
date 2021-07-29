package dtos

import "autoflow/pkg/orm"

type BatchSaveRequest struct {
	ProjectId uint64
	GraphId   uint64

	Nodes       []*orm.Node
	Cards       []*orm.EventCard
	Connections []*orm.Connection
}

type BatchSaveResponse struct {
	ProjectId uint64
	GraphId   uint64

	Nodes       []*orm.Node
	Cards       []*orm.EventCard
	Connections []*orm.Connection
}

type BatchDeleteRequest struct {
	ProjectId   uint64
	GraphId     uint64
	Nodes       []uint64
	Cards       []uint64
	Connections []uint64
}

type BatchDeleteResponse struct {
}
