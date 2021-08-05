package orm

import "gorm.io/datatypes"

type EventCard struct {
	//Id
	ProjectId uint64
	Id        uint64 `gorm:"primarykey"`
	GraphId   uint64

	//Placement
	SourceLocalId uint64
	SlidePort     string

	//Event search
	HttpVote uint64

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

	//UI
	Ui   datatypes.JSON `gorm:"default:null"`
	Name string
}
