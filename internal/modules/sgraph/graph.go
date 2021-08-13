package sgraph

import (
	"autoflow/internal/modules/srepo"
	"autoflow/pkg/common"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/storage"
	"context"
	"gorm.io/gorm"
)

type Graph struct {
	db   *gorm.DB
	repo *srepo.Repo
}

func NewGraph(
	db *gorm.DB,
	repo *srepo.Repo,
) *Graph {
	svc := &Graph{
		db, repo,
	}

	return svc
}

func (s *Graph) SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error) {
	entity := &graph.DBGraph{}
	err := s.repo.SaveProjectObject(c, data, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Graph) SaveNode(c context.Context, node *graph.DBNode) (*graph.DBNode, error) {
	entity := &graph.DBNode{}
	err := s.repo.SaveGraphObject(c, node, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Graph) SaveEventCard(c context.Context, card *graph.DBEventCard) (*graph.DBEventCard, error) {
	entity := &graph.DBEventCard{}
	err := s.repo.SaveGraphObject(c, card, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Graph) SaveConnection(c context.Context, connection *graph.DBConnection) (*graph.DBConnection, error) {
	entity := &graph.DBConnection{}
	err := s.repo.SaveGraphObject(c, connection, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Graph) DeleteGraph(c context.Context, request *common.ProjectModel) error {
	return s.repo.DeleteProjectObject(c, request, &graph.DBGraph{})
}

func (s *Graph) DeleteNode(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBNode{})
}

func (s *Graph) DeleteEventCard(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBEventCard{})
}

func (s *Graph) DeleteConnection(c context.Context, request *graph.IDGraph) error {
	return s.repo.DeleteGraphObject(c, request, &graph.DBConnection{})
}

func (s *Graph) GetGraph(c context.Context, r *common.ProjectModel) (*graph.DBGraph, error) {
	g := &graph.DBGraph{}
	err := s.repo.GetProjectObject(c, r, g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *Graph) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
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
