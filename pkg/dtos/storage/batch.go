package storage

import "autoflow/pkg/orm"

type RequestBatchSave struct {
	ProjectId uint64
	GraphId   uint64

	Nodes       []*orm.Node
	Cards       []*orm.EventCard
	Connections []*orm.Connection
}

type ResponseBatchSave struct {
	ProjectId uint64
	GraphId   uint64

	Nodes       []*orm.Node
	Cards       []*orm.EventCard
	Connections []*orm.Connection
}

type RequestBatchDelete struct {
	ProjectId   uint64
	GraphId     uint64
	Nodes       []uint64
	Cards       []uint64
	Connections []uint64
}

type ResponseBatchDelete struct {
}
