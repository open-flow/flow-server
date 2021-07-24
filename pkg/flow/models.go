package flow

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Graph struct {
	Model

	ProjectID uint64
	UI        datatypes.JSON

	Name string

	Nodes       []*Node
	Events      []*Event
	Connections []*Connection
}

type Event struct {
	Model

	ProjectID    uint64 `gorm:"index:graph_local,unique,priority=1"`
	GraphID      uint64 `gorm:"index:graph_local,priority=2"`
	GraphLocalID uint64 `gorm:"index:graph_local,priority=3"`

	UI datatypes.JSON

	Name string

	Cards []*EventCard
}

type EventCard struct {
	ID        uint64 `gorm:"primarykey"`
	ProjectID uint64
	EventID   uint64

	Platform string `gorm:"index:owner,priority=1"`

	OwnerType string `gorm:"index:owner,priority=2"`
	OwnerID   string `gorm:"index:owner,priority=3"`

	ResourceType string `gorm:"index:resource,priority=1"`
	ResourceID   string `gorm:"index:resource,priority=2"`

	ContextType string `gorm:"index:context,priority=1"`
	ContextID   string `gorm:"index:context,priority=2"`

	InitiatorType string `gorm:"index:initiator,priority=1"`
	InitiatorID   string `gorm:"index:initiator,priority=2"`
}

type Node struct {
	Model

	ProjectID    uint64 `gorm:"index:graph_local,unique,priority=1"`
	GraphID      uint64 `gorm:"index:graph_local,priority=2"`
	GraphLocalID uint64 `gorm:"index:graph_local,priority=3"`

	UI datatypes.JSON

	Name string

	Module   string
	Function string

	Arguments datatypes.JSON
}

type Connection struct {
	ID        uint64 `gorm:"primarykey"`
	ProjectID uint64
	GraphID   uint64

	SourcePort string
	SourceID   uint64

	TargetPort string
	TargetID   uint64
}

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
