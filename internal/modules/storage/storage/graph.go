package storage

import (
	"autoflow/internal/modules/storage/repo"
	"autoflow/pkg/common"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/storage"
	"context"
	"gorm.io/gorm"
)

type Service struct {
	db   *gorm.DB
	repo *repo.Service
}

func New(
	db *gorm.DB,
	repo *repo.Service,
) *Service {
	svc := &Service{
		db, repo,
	}

	return svc
}

func (s *Service) SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error) {
	entity := &graph.DBGraph{}
	err := s.repo.SaveProjectObject(c, data, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service) SaveNode(c context.Context, node *graph.DBNode) (*graph.DBNode, error) {
	entity := &graph.DBNode{}
	err := s.repo.SaveGraphObject(c, node, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service) SaveEventCard(c context.Context, card *graph.DBEventCard) (*graph.DBEventCard, error) {
	entity := &graph.DBEventCard{}
	err := s.repo.SaveGraphObject(c, card, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service) SaveConnection(c context.Context, connection *graph.DBConnection) (*graph.DBConnection, error) {
	entity := &graph.DBConnection{}
	err := s.repo.SaveGraphObject(c, connection, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service) DeleteGraph(c context.Context, request *common.IDProject) error {
	return s.repo.DeleteProjectObject(c, request, &graph.DBGraph{})
}

func (s *Service) DeleteNode(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBNode{})
}

func (s *Service) DeleteEventCard(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBEventCard{})
}

func (s *Service) DeleteConnection(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBConnection{})
}

func (s *Service) GetGraph(c context.Context, r *common.IDProject) (*graph.DBGraph, error) {
	g := &graph.DBGraph{}
	err := s.repo.GetProjectObject(c, r, g)
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
