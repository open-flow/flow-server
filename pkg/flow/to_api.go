package flow

import api "gitlab.com/yautoflow/protorepo-flow-server-go"

func (g *Graph) toAPI() *api.Graph {
	var graph api.Graph

	graph.ProjectID = g.ProjectID
	graph.Name = g.Name
	graph.ID = g.ID
	graph.UI = g.UI.String()

	graph.Nodes = make([]*api.Node, len(g.Nodes))

	for i, n := range g.Nodes {
		graph.Nodes[i] = n.toAPI()
	}

	graph.Events = make([]*api.Event, len(g.Events))

	for i, e := range g.Events {
		graph.Events[i] = e.toAPI()
	}

	return &graph
}

func (n *Node) toAPI() *api.Node {
	var node api.Node

	node.ProjectID = n.ProjectID
	node.ID = n.ID
	node.GraphID = n.GraphID
	node.UI = n.UI.String()
	node.GraphLocalID = n.GraphLocalID
	node.Name = n.Name
	node.Module = n.Module
	node.Function = n.Function
	node.Arguments = n.Arguments.String()

	return &node
}

func (e *Event) toAPI() *api.Event {
	var event api.Event

	event.ProjectID = e.ProjectID
	event.ID = e.ID
	event.GraphID = e.GraphID
	event.GraphLocalID = e.GraphLocalID
	event.Name = e.Name

	event.Cards = make([]*api.EventCard, len(e.Cards))

	for i, c := range e.Cards {
		event.Cards[i] = c.toAPI()
	}

	return &event
}

func (c *EventCard) toAPI() *api.EventCard {
	var card api.EventCard

	card.ID = c.ID
	card.ProjectID = c.ProjectID
	card.EventID = c.EventID
	card.Platform = c.Platform
	card.OwnerType = c.OwnerType
	card.OwnerID = c.OwnerID
	card.ResourceType = c.ResourceType
	card.ResourceID = c.ResourceID
	card.ContextType = c.ContextType
	card.ContextID = c.ContextID
	card.InitiatorType = c.ContextType
	card.InitiatorID = c.InitiatorID

	return &card
}
