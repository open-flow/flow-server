package graph

func (g *DBGraph) FindNode(localId uint64) DBNode {
	for _, n := range g.Nodes {
		if n.LocalId == localId {
			return n
		}
	}

	return DBNode{}
}

func (g *DBGraph) FindConnectedNodes(localId uint64, slidePort string) []DataConnection {
	if localId == 0 || slidePort == "" {
		return nil
	}

	var nodes []DataConnection

	for _, c := range g.Connections {
		if c.SourceId == localId && c.SourcePort == slidePort {
			nodes = append(nodes, DataConnection{
				SourcePort: c.SourcePort,
				SourceId:   c.SourceId,
				TargetPort: c.TargetPort,
				TargetId:   c.TargetId,
			})
		}
	}

	return nodes
}
