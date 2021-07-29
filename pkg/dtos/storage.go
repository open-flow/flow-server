package dtos

import "autoflow/pkg/orm"

type DeleteRequest struct {
	ProjectId uint64
	Id        uint64
}

type DeleteResponse struct {
}

type GetFullGraphRequest struct {
	ProjectId uint64
	Id        uint64
}

type ListGraphRequest struct {
	ProjectIds []uint64
}

type ListGraphResponse struct {
	Graphs []*orm.Graph
}
