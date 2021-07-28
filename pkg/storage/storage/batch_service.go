package storage

import (
	"context"
	"gitlab.com/yautoflow/interfaces/dtos"
	"gitlab.com/yautoflow/interfaces/orm"
	"gorm.io/gorm"
)

type BatchService struct {
	db *gorm.DB
}

func NewBatchService(db *gorm.DB) *BatchService {
	return &BatchService{
		db: db,
	}
}

func (s *BatchService) Save(ctx context.Context, r *dtos.BatchSaveRequest) (*dtos.BatchSaveResponse, error) {
	graph := &orm.Graph{}

	for _, v := range graph.Nodes {
		v.ProjectId = r.ProjectId
		v.GraphId = r.GraphId
		v.Id = 0
	}

	for _, v := range graph.Cards {
		v.ProjectId = r.ProjectId
		v.GraphId = r.GraphId
		v.Id = 0
	}

	for _, v := range graph.Connections {
		v.ProjectId = r.ProjectId
		v.GraphId = r.GraphId
		v.Id = 0
	}

	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_Id = ? and Id = ?", r.ProjectId, r.GraphId).
				First(graph)

			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(graph.Nodes)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(graph.Cards)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(graph.Connections)
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &dtos.BatchSaveResponse{
		ProjectId:   r.ProjectId,
		GraphId:     r.GraphId,
		Nodes:       graph.Nodes,
		Cards:       graph.Cards,
		Connections: graph.Connections,
	}, nil
}

func (s *BatchService) Delete(ctx context.Context, r *dtos.BatchDeleteRequest) (*dtos.BatchDeleteResponse, error) {
	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"project_Id = ? and graph_Id = ? and Id in ?",
				r.ProjectId,
				r.GraphId,
				r.Connections,
			).Delete(&orm.Connection{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_Id = ? and graph_Id = ? and Id in ?",
					r.ProjectId,
					r.GraphId,
					r.Cards,
				).Delete(&orm.EventCard{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_Id = ? and graph_Id = ? and Id in ?",
					r.ProjectId,
					r.GraphId,
					r.Nodes,
				).Delete(&orm.Node{})
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &dtos.BatchDeleteResponse{}, nil
}
