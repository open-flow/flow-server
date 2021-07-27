package flow

import (
	"context"
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/flow-proto/gen/go/flow/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type graphService struct {
	db *gorm.DB
	api.UnimplementedGraphServiceServer
}

func NewGraphService(db *gorm.DB) api.GraphServiceServer {
	return &graphService{
		db: db,
	}
}

func (s *graphService) GetDB() *gorm.DB {
	return s.db
}

func (s *graphService) SaveGraph(c context.Context, graph *api.Graph) (*api.Graph, error) {
	err := StoreGeneric(s, c, graph, &Graph{})
	if err != nil {
		return nil, err
	}
	return graph, nil
}

func (s *graphService) SaveNode(c context.Context, node *api.Node) (*api.Node, error) {
	err := StoreGeneric(s, c, node, &Node{})
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *graphService) SaveEventCard(c context.Context, card *api.EventCard) (*api.EventCard, error) {
	err := StoreGeneric(s, c, card, &EventCard{})
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *graphService) SaveConnection(c context.Context, connection *api.Connection) (*api.Connection, error) {
	err := StoreGeneric(s, c, connection, &Connection{})
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (s *graphService) DeleteGraph(c context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	return DeleteGeneric(s, c, request, &Graph{})
}

func (s *graphService) DeleteNode(c context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	return DeleteGeneric(s, c, request, &Node{})
}

func (s *graphService) DeleteEventCard(c context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	return DeleteGeneric(s, c, request, &EventCard{})
}

func (s *graphService) DeleteConnection(c context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	return DeleteGeneric(s, c, request, &Connection{})
}

func (s *graphService) GetFullGraph(c context.Context, request *api.GetFullGraphRequest) (*api.GetFullGraphResponse, error) {
	var graph = &Graph{}
	var deepGraph = &api.GetFullGraphResponse{
		Graph: &api.Graph{},
	}

	err := s.GetDB().
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", request.GetProjectId(), request.GetId()).
				Preload(clause.Associations).
				First(graph)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(deepGraph.Graph, graph)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&deepGraph.Cards, graph.Cards)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&deepGraph.Nodes, graph.Nodes)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&deepGraph.Connections, graph.Connections)
	if err != nil {
		return nil, err
	}

	return deepGraph, nil
}

func (s *graphService) ListGraph(c context.Context, request *api.ListGraphRequest) (*api.ListGraphResponse, error) {
	list := &api.ListGraphResponse{}
	var graphs []*Graph

	err := s.GetDB().
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id in ?", request.ProjectIds).
				Find(&graphs)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&list.Graphs, graphs)
	if err != nil {
		return nil, err
	}

	return list, nil
}
