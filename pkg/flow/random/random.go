package random

import (
	"autoflow/pkg/flow"
	gofakeit "github.com/brianvoe/gofakeit/v6"
	"gorm.io/datatypes"
)

func Graph() *flow.Graph {
	fake := gofakeit.NewCrypto()

	var nodesCount = fake.Number(100, 200)
	var cardsCount = fake.Number(10, 50)
	var connectionsCount = fake.Number(100, 200)

	var graph flow.Graph
	graph.Nodes = make([]*flow.Node, nodesCount)
	for i := 0; i < nodesCount; i++ {
		graph.Nodes[i] = &flow.Node{
			GraphId:   0,
			ProjectId: 1,
			LocalId:   uint64(i),
			Arguments: datatypes.JSON("{}"),
			Ui:        datatypes.JSON("{}"),

			Name:     fake.PetName(),
			Module:   fake.AppName(),
			Function: fake.URL(),
		}
	}

	graph.Cards = make([]*flow.EventCard, cardsCount)
	for i := 0; i < cardsCount; i++ {
		var event = &flow.EventCard{
			GraphId:    0,
			ProjectId:  1,
			TargetId:   uint64(fake.Number(0, nodesCount)),
			Platform:   "random",
			StaticType: "random",
			StaticId:   fake.UUID(),
		}

		graph.Cards[i] = event
	}

	graph.Connections = make([]*flow.Connection, connectionsCount)
	for i := 0; i < connectionsCount; i++ {
		graph.Connections[i] = &flow.Connection{
			GraphId:   0,
			ProjectId: 1,

			SourcePort: fake.Name(),
			SourceId:   fake.Uint64(),

			TargetPort: fake.Name(),
			TargetId:   fake.Uint64(),
		}
	}

	graph.Name = fake.AppName()
	graph.ProjectId = 1

	return &graph
}
