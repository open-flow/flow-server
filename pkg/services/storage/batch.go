package storage

import (
	"autoflow/pkg/dtos/storage"
	"autoflow/pkg/orm"
	"context"
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

func (s *BatchService) Save(ctx context.Context, r *storage.RequestBatchSave) (*storage.ResponseBatchSave, error) {
	graph := &orm.Graph{}

	for _, v := range r.Nodes {
		v.ProjectId = r.ProjectId
		v.GraphId = r.GraphId
		v.Id = 0
	}

	for _, v := range r.Cards {
		v.ProjectId = r.ProjectId
		v.GraphId = r.GraphId
		v.Id = 0
	}

	for _, v := range r.Connections {
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
				Where("project_id = ? and id = ?", r.ProjectId, r.GraphId).
				First(graph)

			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(r.Nodes)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(r.Cards)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(r.Connections)
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &storage.ResponseBatchSave{
		ProjectId:   r.ProjectId,
		GraphId:     r.GraphId,
		Nodes:       r.Nodes,
		Cards:       r.Cards,
		Connections: r.Connections,
	}, nil
}

func (s *BatchService) Delete(ctx context.Context, r *storage.RequestBatchDelete) (*storage.ResponseBatchDelete, error) {
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

	return &storage.ResponseBatchDelete{}, nil
}
