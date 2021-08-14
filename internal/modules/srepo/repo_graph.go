package srepo

import (
	"autoflow/pkg/storage/graph"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Repo) SaveGraphObject(ctx context.Context, save graph.Object, entity graph.Object) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and graph_id = ? and id = ?",
				save.GetProjectId(),
				save.GetGraphId(),
				save.GetId(),
			).
				Assign(save).
				FirstOrCreate(entity)
			if r.Error != nil {
				return r.Error
			}
			return nil
		})

	if err != nil {
		s.logger.Error("saving graph object failed", zap.Error(err), zap.Any("save", save))
		return err
	}

	return nil
}

func (s *Repo) GetGraphObject(ctx context.Context, object graph.Object, entity graph.Object) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and graph_id = ? and id = ?",
				object.GetProjectId(),
				object.GetGraphId(),
				object.GetId(),
			).First(entity)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})

	if err != nil {
		s.logger.Error("getting graph object failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Repo) DeleteGraphObject(ctx context.Context, id *graph.GraphObject, entity graph.Object) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and graph_id = ? and id = ?",
				id.GetProjectId(),
				id.GetGraphId(),
				id.GetId(),
			).Delete(entity)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		s.logger.Error("deleting graph object failed", zap.Error(err))
		return err
	}

	return nil
}
