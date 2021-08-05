package utils

import (
	"autoflow/pkg/orm"
	gofakeit "github.com/brianvoe/gofakeit/v6"
)

func Graph() *orm.Graph {
	fake := gofakeit.NewCrypto()

	var nodesCount = fake.Number(100, 200)
	var cardsCount = fake.Number(10, 50)
	var connectionsCount = fake.Number(100, 200)

	var graph orm.Graph
	graph.Nodes = make([]*orm.Node, nodesCount)
	for i := 0; i < nodesCount; i++ {
		graph.Nodes[i] = &orm.Node{
			GraphId:   0,
			ProjectId: 1,
			LocalId:   uint64(i),
			Arguments: "{}",
			Ui:        "{}",

			Name:     fake.PetName(),
			Module:   fake.AppName(),
			Function: fake.URL(),
		}
	}

	graph.Cards = make([]*orm.EventCard, cardsCount)
	for i := 0; i < cardsCount; i++ {
		var event = &orm.EventCard{
			GraphId:    0,
			ProjectId:  1,
			TargetId:   uint64(fake.Number(0, nodesCount)),
			Platform:   "random",
			StaticType: "random",
			StaticId:   fake.UUID(),
		}

		graph.Cards[i] = event
	}

	graph.Connections = make([]*orm.Connection, connectionsCount)
	for i := 0; i < connectionsCount; i++ {
		graph.Connections[i] = &orm.Connection{
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
