package flow

import (
	"bytes"
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
)

func (g *Graph) assignID(graph *api.Graph) {
	g.ID = graph.ID
	g.ProjectID = graph.ProjectID
}

func (n *Node) assignID(node *api.Node) {
	n.ID = node.ID
	n.ProjectID = node.ProjectID
	n.GraphID = node.GraphID
}

func (e *Event) assignID(event *api.Event) {
	e.ID = event.ID
	e.ProjectID = event.ProjectID
	e.GraphID = event.GraphID
}

func (c *EventCard) assignID(card *api.EventCard) {
	c.ID = card.ID
	c.ProjectID = card.ProjectID
}

func (c *Connection) assignID(card *api.EventCard) {
	c.ID = card.ID
	c.ProjectID = card.ProjectID
}

func (g *Graph) assign(source *api.Graph) bool {
	var changed bool

	ui := []byte(source.UI)
	if !bytes.Equal(g.UI, ui) {
		g.UI = ui
		changed = true
	}

	if g.Name != source.Name {
		g.Name = source.Name
		changed = true
	}

	return changed
}

func (n *Node) assign(source *api.Node) bool {
	var changed bool

	ui := []byte(source.UI)
	if !bytes.Equal(n.UI, ui) {
		n.UI = ui
		changed = true
	}

	if n.GraphLocalID != source.GraphLocalID {
		n.GraphLocalID = source.GraphLocalID
		changed = true
	}

	if n.Name != source.Name {
		n.Name = source.Name
		changed = true
	}

	if n.Module != source.Module {
		n.Module = source.Module
		changed = true
	}

	if n.Function != source.Function {
		n.Function = source.Function
		changed = true
	}

	arguments := []byte(source.Arguments)
	if !bytes.Equal(n.Arguments, arguments) {
		n.Arguments = arguments
		changed = true
	}

	return changed
}

func (e *Event) assign(source *api.Event) bool {
	var changed bool

	ui := []byte(source.UI)
	if !bytes.Equal(e.UI, ui) {
		e.UI = ui
		changed = true
	}

	if e.GraphLocalID != source.GraphLocalID {
		e.GraphLocalID = source.GraphLocalID
		changed = true
	}

	if e.Name != source.Name {
		e.Name = source.Name
		changed = true
	}

	return changed
}

func (c *EventCard) assign(source *api.EventCard) bool {
	var changed bool

	if c.EventID != source.EventID {
		c.EventID = source.EventID
		changed = true
	}

	if c.Platform != source.Platform {
		c.Platform = source.Platform
		changed = true
	}

	if c.OwnerType != source.OwnerType {
		c.OwnerType = source.OwnerType
		changed = true
	}

	if c.OwnerID != source.OwnerID {
		c.OwnerID = source.OwnerID
		changed = true
	}

	if c.ResourceType != source.ResourceType {
		c.ResourceType = source.ResourceType
		changed = true
	}

	if c.ResourceID != source.ResourceID {
		c.ResourceID = source.ResourceID
		changed = true
	}

	if c.ContextType != source.ContextType {
		c.ContextType = source.ContextType
		changed = true
	}

	if c.ContextID != source.ContextID {
		c.ContextID = source.ContextID
		changed = true
	}

	if c.InitiatorType != source.InitiatorType {
		c.InitiatorType = source.InitiatorType
		changed = true
	}

	if c.InitiatorID != source.InitiatorID {
		c.InitiatorID = source.InitiatorID
		changed = true
	}

	return changed
}
