package storage

import (
	"autoflow/pkg/storage/dtos"
	"autoflow/pkg/storage/orm"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GraphService struct {
	db *gorm.DB
}

func NewGraphService(db *gorm.DB) *GraphService {
	return &GraphService{
		db: db,
	}
}

func (s *GraphService) SaveGraph(c context.Context, graph *orm.Graph) (*orm.Graph, error) {
	err := s.StoreGeneric(c, graph)
	if err != nil {
		return nil, err
	}
	return graph, nil
}

func (s *GraphService) SaveNode(c context.Context, node *orm.Node) (*orm.Node, error) {
	err := s.StoreGeneric(c, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *GraphService) SaveEventCard(c context.Context, card *orm.EventCard) (*orm.EventCard, error) {
	err := s.StoreGeneric(c, card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *GraphService) SaveConnection(c context.Context, connection *orm.Connection) (*orm.Connection, error) {
	err := s.StoreGeneric(c, connection)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (s *GraphService) DeleteGraph(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Graph{})
}

func (s *GraphService) DeleteNode(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Node{})
}

func (s *GraphService) DeleteEventCard(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.EventCard{})
}

func (s *GraphService) DeleteConnection(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Connection{})
}

func (s *GraphService) GetFullGraph(c context.Context, r *dtos.GetFullGraphRequest) (*orm.Graph, error) {
	var graph = &orm.Graph{}

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", r.ProjectId, r.Id).
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

	return graph, nil
}

func (s *GraphService) ListGraph(c context.Context, r *dtos.ListGraphRequest) (*dtos.ListGraphResponse, error) {
	var graphs []*orm.Graph

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id in ?", r.ProjectIds).
				Find(&graphs)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &dtos.ListGraphResponse{
		Graphs: graphs,
	}, nil
}
