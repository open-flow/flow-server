package flow

import (
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"gorm.io/gorm"
)

type service struct {
	orm gorm.DB
}

func (s *service) ConvertRawGraph(graph *api.Graph) *Graph {
	var entity Graph
	entity.assignID(graph)
	entity.assign(graph)

	entity.Events = make([]*Event, len(graph.Events))
	for i, e := range graph.Events {
		var eventEntity Event
		eventEntity.assignID(e)
		eventEntity.assign(e)

		eventEntity.Cards = make([]*EventCard, len(e.Cards))
		for o, c := range e.Cards {
			var eventCardEntity EventCard
			eventCardEntity.assignID(c)
			eventCardEntity.assign(c)
			eventEntity.Cards[o] = &eventCardEntity
		}

		entity.Events[i] = &eventEntity
	}

	entity.Nodes = make([]*Node, len(graph.Nodes))
	for i, n := range graph.Nodes {
		var nodeEntity Node
		nodeEntity.assignID(n)
		nodeEntity.assign(n)

		entity.Nodes[i] = &nodeEntity
	}
}

func (s *service) SaveGraph(graph Graph) error {

	return nil
}

func (s *service) SaveNode(node Node) error {

	return nil
}

func (s *service) SaveEvent(event Event) error {

	return nil
}

func (s *service) SaveEventCard(card Event) error {

	return nil
}
