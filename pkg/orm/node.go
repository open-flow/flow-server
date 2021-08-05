package orm

import "gorm.io/datatypes"

type Node struct {
	//ID
	ProjectId uint64 `gorm:"index:graph_local,unique,priority=1"`
	Id        uint64 `gorm:"primarykey"`
	GraphId   uint64 `gorm:"index:graph_local,priority=2"`

	//Placement
	LocalId uint64 `gorm:"index:graph_local,priority=3"`

	//Invocation
	Type      string
	Module    string
	Function  string
	Arguments datatypes.JSON `gorm:"default:null"`

	//UI
	Ui   datatypes.JSON `gorm:"default:null"`
	Name string
}
