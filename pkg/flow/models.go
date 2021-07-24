package flow

import (
	"database/sql"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Graph struct {
	gorm.Model

	ProjectID uint
	UI        datatypes.JSON

	Name string

	Nodes  []Node
	Events []Event
}

type Event struct {
	gorm.Model

	ProjectID    uint `gorm:"index:graph_local,unique,priority=1"`
	GraphID      uint `gorm:"index:graph_local,priority=2"`
	GraphLocalID uint `gorm:"index:graph_local,priority=3"`

	UI datatypes.JSON

	Name string

	Cards []EventCard
}

type EventCard struct {
	ID        uint `gorm:"primarykey"`
	ProjectID uint
	EventID   uint

	Platform string `gorm:"index:owner,priority=1"`

	OwnerType string `gorm:"index:owner,priority=2"`
	OwnerID   string `gorm:"index:owner,priority=3"`

	ResourceType sql.NullString `gorm:"index:resource,priority=1"`
	ResourceID   sql.NullString `gorm:"index:resource,priority=2"`

	ContextType sql.NullString `gorm:"index:context,priority=1"`
	ContextID   sql.NullString `gorm:"index:context,priority=2"`

	InitiatorType sql.NullString `gorm:"index:initiator,priority=1"`
	InitiatorID   sql.NullString `gorm:"index:initiator,priority=2"`
}

type Node struct {
	gorm.Model

	ProjectID    uint `gorm:"index:graph_local,unique,priority=1"`
	GraphID      uint `gorm:"index:graph_local,priority=2"`
	GraphLocalID uint `gorm:"index:graph_local,priority=3"`

	UI datatypes.JSON

	Module   string
	Function string

	Arguments datatypes.JSON
}
