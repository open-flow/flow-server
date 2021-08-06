package graph

//
//import "github.com/brianvoe/gofakeit/v6"
//
//func Random() *DBGraph {
//	fake := gofakeit.NewCrypto()
//
//	var nodesCount = fake.Number(100, 200)
//	var cardsCount = fake.Number(10, 50)
//	var connectionsCount = fake.Number(100, 200)
//	var graph DBGraph
//	idGraph := IDGraph{
//		IDProject: IDProject{
//			ProjectID: 1,
//		},
//	}
//
//	graph.Nodes = make([]*DBNode, nodesCount)
//	for i := 0; i < nodesCount; i++ {
//		graph.Nodes[i] = &DBNode{
//			IDGraph: idGraph,
//			DataNode: DataNode{
//				LocalId:  uint64(i),
//				Module:   fake.AppName(),
//				Function: fake.URL(),
//			},
//			DataUI: DataUI{
//				Name: fake.PetName(),
//			},
//		}
//	}
//
//	graph.Cards = make([]*DBEventCard, cardsCount)
//	for i := 0; i < cardsCount; i++ {
//		var event = &DBEventCard{
//			IDGraph: idGraph,
//
//			DataEventCard: DataEventCard{
//				TargetId: uint64(fake.Number(0, nodesCount)),
//
//				DataEvent: DataEvent{
//					Platform:   "random",
//					StaticType: "random",
//					StaticId:   fake.UUID(),
//				},
//			},
//		}
//
//		graph.Cards[i] = event
//	}
//
//	graph.Connections = make([]*DBConnection, connectionsCount)
//	for i := 0; i < connectionsCount; i++ {
//		graph.Connections[i] = &DBConnection{
//			IDGraph: idGraph,
//
//			DataConnection: DataConnection{
//				SourcePort: fake.Name(),
//				SourceId:   fake.Uint64(),
//
//				TargetPort: fake.Name(),
//				TargetId:   fake.Uint64(),
//			},
//		}
//	}
//
//	graph.Name = fake.AppName()
//	graph.ProjectID = 1
//
//	return &graph
//}
