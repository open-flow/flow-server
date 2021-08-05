package orm

import (
	"autoflow/pkg/dtos/execution"
	"gorm.io/datatypes"
)

type Graph struct {
	//ID
	ProjectId uint64
	Id        uint64 `gorm:"primarykey"`

	//Params
	LocalIdCounter uint64

	//UI
	Ui   datatypes.JSON `gorm:"default:null"`
	Name string

	//Relations
	Nodes       []*Node       `gorm:"constraint:OnDelete:CASCADE;"`
	Cards       []*EventCard  `gorm:"constraint:OnDelete:CASCADE;"`
	Connections []*Connection `gorm:"constraint:OnDelete:CASCADE;"`
}

func (g *Graph) FindNode(localId uint64) *Node {
	for _, n := range g.Nodes {
		if n.LocalId == localId {
			return n
		}
	}

	return nil
}

func (g *Graph) FindConnectedNodes(localId uint64, slidePort string) []*execution.Connection {
	if localId == 0 || slidePort == "" {
		return nil
	}

	var nodes []*execution.Connection

	for _, c := range g.Connections {
		if c.SourceId == localId && c.SourcePort == slidePort {
			nodes = append(nodes, &execution.Connection{
				SourcePort: c.SourcePort,
				SourceId:   c.SourceId,
				TargetPort: c.TargetPort,
				TargetId:   c.TargetId,
			})
		}
	}

	return nodes
}
