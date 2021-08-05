package data

import (
	"autoflow/pkg/dtos/execution"
	"autoflow/pkg/orm"
)

func FindNode(g *orm.Graph, localId uint64) *orm.Node {
	for _, n := range g.Nodes {
		if n.LocalId == localId {
			return n
		}
	}

	return nil
}

func FindConnectedNodes(g *orm.Graph, localId uint64, slidePort string) []*execution.Connection {
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
