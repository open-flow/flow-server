package random

import (
	"autoflow/pkg/flow"
	gofakeit "github.com/brianvoe/gofakeit/v6"
	"gorm.io/datatypes"
)

func Graph() *flow.Graph {
	fake := gofakeit.NewCrypto()

	var nodesCount = fake.Number(100, 200)
	var eventsCount = fake.Number(10, 50)
	var connectionsCount = fake.Number(100, 200)

	var graph flow.Graph
	graph.Nodes = make([]*flow.Node, nodesCount)
	for i := 0; i < nodesCount; i++ {
		graph.Nodes[i] = &flow.Node{
			GraphID:      0,
			ProjectID:    1,
			GraphLocalID: uint64(i),
			Arguments:    datatypes.JSON("{}"),
			UI:           datatypes.JSON("{}"),

			Name:     fake.PetName(),
			Module:   fake.AppName(),
			Function: fake.URL(),
		}
	}

	graph.Events = make([]*flow.Event, eventsCount)
	for i := 0; i < eventsCount; i++ {
		var cardsCount = fake.Number(1, 5)
		var event = &flow.Event{
			GraphID:      0,
			ProjectID:    1,
			GraphLocalID: uint64(i),
			UI:           datatypes.JSON("{}"),

			Name:  fake.PetName(),
			Cards: make([]*flow.EventCard, cardsCount),
		}

		for o := 0; o < cardsCount; o++ {
			event.Cards[o] = &flow.EventCard{
				ID:        0,
				ProjectID: 1,
				Platform:  fake.AppName(),
				OwnerType: fake.PetName(),
				OwnerID:   fake.UUID(),
			}
		}

		graph.Events[i] = event
	}

	graph.Connections = make([]*flow.Connection, connectionsCount)
	for i := 0; i < connectionsCount; i++ {
		graph.Connections[i] = &flow.Connection{
			GraphID:   0,
			ProjectID: 1,

			SourcePort: fake.Name(),
			SourceID:   fake.Uint64(),

			TargetPort: fake.Name(),
			TargetID:   fake.Uint64(),
		}
	}

	graph.Name = fake.AppName()
	graph.ProjectID = 1

	return &graph
}
