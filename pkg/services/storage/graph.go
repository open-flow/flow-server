package storage

import (
	"autoflow/pkg/dtos/storage"
	"autoflow/pkg/orm"
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
	var persisted orm.Graph
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", graph.ProjectId, graph.Id).
				Assign(graph).
				FirstOrCreate(&persisted)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &persisted, nil
}

func (s *GraphService) SaveNode(c context.Context, node *orm.Node) (*orm.Node, error) {
	var persisted orm.Node

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", node.ProjectId, node.Id).
				Assign(node).
				FirstOrCreate(&persisted)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &persisted, nil
}

func (s *GraphService) SaveEventCard(c context.Context, card *orm.EventCard) (*orm.EventCard, error) {
	var persisted orm.EventCard

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", card.ProjectId, card.Id).
				Assign(card).
				FirstOrCreate(&persisted)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &persisted, nil
}

func (s *GraphService) SaveConnection(c context.Context, connection *orm.Connection) (*orm.Connection, error) {
	var persisted orm.Connection

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", connection.ProjectId, connection.Id).
				Assign(connection).
				FirstOrCreate(&persisted)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &persisted, nil
}

func (s *GraphService) DeleteGraph(c context.Context, request *storage.RequestDelete) (*storage.ResponseDelete, error) {
	return s.DeleteGeneric(c, request, &orm.Graph{})
}

func (s *GraphService) DeleteNode(c context.Context, request *storage.RequestDelete) (*storage.ResponseDelete, error) {
	return s.DeleteGeneric(c, request, &orm.Node{})
}

func (s *GraphService) DeleteEventCard(c context.Context, request *storage.RequestDelete) (*storage.ResponseDelete, error) {
	return s.DeleteGeneric(c, request, &orm.EventCard{})
}

func (s *GraphService) DeleteConnection(c context.Context, request *storage.RequestDelete) (*storage.ResponseDelete, error) {
	return s.DeleteGeneric(c, request, &orm.Connection{})
}

func (s *GraphService) GetFullGraph(c context.Context, r *storage.RequestGetFullGraph) (*orm.Graph, error) {
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

func (s *GraphService) ListGraph(c context.Context, r *storage.RequestListGraph) (*storage.ResponseListGraph, error) {
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

	return &storage.ResponseListGraph{
		Graphs: graphs,
	}, nil
}

func (s *GraphService) StoreGeneric(c context.Context, model interface{}) error {
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			//todo change to use projectID
			res := tx.FirstOrCreate(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return err
	}

	return nil
}

func (s *GraphService) DeleteGeneric(c context.Context, req *storage.RequestDelete, model interface{}) (*storage.ResponseDelete, error) {
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", req.ProjectId, req.Id).
				Delete(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &storage.ResponseDelete{}, nil
}
