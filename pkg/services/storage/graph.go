package storage

import (
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/storage"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error) {
	var persisted graph.DBGraph
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", data.ProjectId, data.Id).
				Assign(data).
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

func (s *Service) SaveNode(c context.Context, node *graph.DBNode) (*graph.DBNode, error) {
	var persisted graph.DBNode

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

func (s *Service) SaveEventCard(c context.Context, card *graph.DBEventCard) (*graph.DBEventCard, error) {
	var persisted graph.DBEventCard

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

func (s *Service) SaveConnection(c context.Context, connection *graph.DBConnection) (*graph.DBConnection, error) {
	var persisted graph.DBConnection

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

func (s *Service) DeleteGraph(c context.Context, request *storage.DeleteRequest) (*storage.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &graph.DBGraph{})
}

func (s *Service) DeleteNode(c context.Context, request *storage.DeleteRequest) (*storage.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &graph.DBNode{})
}

func (s *Service) DeleteEventCard(c context.Context, request *storage.DeleteRequest) (*storage.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &graph.DBEventCard{})
}

func (s *Service) DeleteConnection(c context.Context, request *storage.DeleteRequest) (*storage.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &graph.DBConnection{})
}

func (s *Service) GetGraph(c context.Context, r *storage.GetGraphRequest) (*graph.DBGraph, error) {
	var g = &graph.DBGraph{}

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", r.ProjectId, r.Id).
				Preload(clause.Associations).
				First(g)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (s *Service) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
	var graphs []graph.DBGraph

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id in ?", r.ProjectId).
				Find(&graphs)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &storage.ListGraphResponse{
		Graphs: graphs,
	}, nil
}

func (s *Service) DeleteGeneric(c context.Context, req *storage.DeleteRequest, model interface{}) (*storage.DeleteResponse, error) {
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

	return &storage.DeleteResponse{}, nil
}
