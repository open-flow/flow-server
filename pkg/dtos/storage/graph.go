package storage

import "autoflow/pkg/orm"

type RequestDelete struct {
	ProjectId uint64
	Id        uint64
}

type ResponseDelete struct {
}

type RequestGetFullGraph struct {
	ProjectId uint64
	Id        uint64
}

type RequestListGraph struct {
	ProjectIds []uint64
}

type ResponseListGraph struct {
	Graphs []*orm.Graph
}
