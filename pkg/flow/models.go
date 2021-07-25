package flow

import (
	"gorm.io/datatypes"
)

type Graph struct {
	ProjectID uint64

	ID uint64 `gorm:"primarykey"`

	UI   datatypes.JSON `gorm:"default:null"`
	Name string

	Nodes       []*Node       `gorm:"constraint:OnDelete:CASCADE;"`
	Cards       []*EventCard  `gorm:"constraint:OnDelete:CASCADE;"`
	Connections []*Connection `gorm:"constraint:OnDelete:CASCADE;"`
}

type Node struct {
	ProjectID uint64 `gorm:"index:graph_local,unique,priority=1"`

	ID      uint64 `gorm:"primarykey"`
	GraphID uint64 `gorm:"index:graph_local,priority=2"`
	LocalID uint64 `gorm:"index:graph_local,priority=3"`

	UI        datatypes.JSON `gorm:"default:null"`
	Name      string
	Type      string
	Module    string
	Function  string
	Arguments datatypes.JSON `gorm:"default:null"`
}

type EventCard struct {
	ProjectID uint64

	ID      uint64 `gorm:"primarykey"`
	GraphID uint64

	TargetID uint64

	UI datatypes.JSON `gorm:"default:null"`

	Platform string `gorm:"index:owner,priority=1"`

	OwnerType string `gorm:"index:owner,priority=2"`
	OwnerID   string `gorm:"index:owner,priority=3"`

	ResourceType string `gorm:"index:resource,priority=1"`
	ResourceID   string `gorm:"index:resource,priority=2"`

	ContextType string `gorm:"index:context,priority=1"`
	ContextID   string `gorm:"index:context,priority=2"`

	InitiatorType string `gorm:"index:initiator,priority=1"`
	InitiatorID   string `gorm:"index:initiator,priority=2"`

	StaticType string `gorm:"index:static,priority=1"`
	StaticID   string `gorm:"index:static,priority=2"`
}

type Connection struct {
	ProjectID uint64

	ID      uint64 `gorm:"primarykey"`
	GraphID uint64

	SourcePort string
	SourceID   uint64

	TargetPort string
	TargetID   uint64
}
