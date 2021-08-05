package execution

import (
	"autoflow/pkg/orm"
)

type State struct {
	Graph      *orm.Graph
	Activation *Activation
	Cursor     *Cursor
	Memory     *Memory
}

type Activation struct {
	Card       *orm.EventCard
	RawRequest interface{}
}

type Cursor struct {
	Node *orm.Node

	Current *Connection
	Path    []*Connection
	Next    []*Connection
}

type Memory struct {
	Context  map[string]interface{}
	Response interface{}
}

type Connection struct {
	SourcePort string
	SourceId   uint64

	TargetPort string
	TargetId   uint64
}

func (c *Connection) Copy() *Connection {
	return &Connection{
		SourcePort: c.SourcePort,
		SourceId:   c.SourceId,
		TargetPort: c.TargetPort,
		TargetId:   c.TargetId,
	}
}
