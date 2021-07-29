package orm

import (
	"gorm.io/datatypes"
)

type Graph struct {
	ProjectId uint64

	Id uint64 `gorm:"primarykey"`

	Ui   datatypes.JSON `gorm:"default:null"`
	Name string

	Nodes       []*Node       `gorm:"constraint:OnDelete:CASCADE;"`
	Cards       []*EventCard  `gorm:"constraint:OnDelete:CASCADE;"`
	Connections []*Connection `gorm:"constraint:OnDelete:CASCADE;"`
}

type Node struct {
	ProjectId uint64 `gorm:"index:graph_local,unique,priority=1"`

	Id      uint64 `gorm:"primarykey"`
	GraphId uint64 `gorm:"index:graph_local,priority=2"`
	LocalId uint64 `gorm:"index:graph_local,priority=3"`

	Ui        datatypes.JSON `gorm:"default:null"`
	Name      string
	Type      string
	Module    string
	Function  string
	Arguments datatypes.JSON `gorm:"default:null"`
}

type EventCard struct {
	ProjectId uint64

	Id      uint64 `gorm:"primarykey"`
	GraphId uint64

	TargetId uint64

	Ui datatypes.JSON `gorm:"default:null"`

	Platform string `gorm:"index:owner,priority=1"`

	OwnerType string `gorm:"index:owner,priority=2"`
	OwnerId   string `gorm:"index:owner,priority=3"`

	ResourceType string `gorm:"index:resource,priority=1"`
	ResourceId   string `gorm:"index:resource,priority=2"`

	ContextType string `gorm:"index:context,priority=1"`
	ContextId   string `gorm:"index:context,priority=2"`

	InitiatorType string `gorm:"index:initiator,priority=1"`
	InitiatorId   string `gorm:"index:initiator,priority=2"`

	StaticType string `gorm:"index:static,priority=1"`
	StaticId   string `gorm:"index:static,priority=2"`
}

type Connection struct {
	ProjectId uint64

	Id      uint64 `gorm:"primarykey"`
	GraphId uint64

	SourcePort string
	SourceId   uint64

	TargetPort string
	TargetId   uint64
}
