package graph

import "autoflow/pkg/entities/common"

type IDGraph struct {
	common.IDProject
	GraphId uint `gorm:"index" json:"graphId,omitempty"`
}

func (i *IDGraph) GetGraphId() uint {
	return i.GraphId
}

type DBGraph struct {
	common.IDProject

	common.DataUI
	DataGraph

	Nodes       []DBNode       `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"nodes,omitempty"`
	Cards       []DBEventCard  `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"cards,omitempty"`
	Connections []DBConnection `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"connections,omitempty"`
}

type DataGraph struct {
	Counter uint `json:"counter"`
}

type Object interface {
	common.ProjectObject
	GetGraphId() uint
}

var _ Object = (*IDGraph)(nil)
