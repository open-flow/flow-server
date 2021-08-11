package repo

import (
	"autoflow/pkg/common"
	"autoflow/pkg/storage/graph"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Service) SaveProjectObject(ctx context.Context, save common.ProjectObject, entity common.ProjectObject) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				save.GetProjectId(),
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
		s.logger.Error("saving project object failed", zap.Error(err), zap.Any("save", save))
		return err
	}

	return nil
}

func (s *Service) SaveGraphObject(ctx context.Context, save graph.Object, entity graph.Object) error {
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

func (s *Service) GetGraphObject(ctx context.Context, object graph.Object, entity graph.Object) error {
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

func (s *Service) GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				object.GetProjectId(),
				object.GetId(),
			).First(entity)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})

	if err != nil {
		s.logger.Error("getting project object failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) DeleteProjectObject(ctx context.Context, id *common.IDProject, entity common.ProjectObject) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				id.GetProjectId(),
				id.GetId(),
			).Delete(entity)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		s.logger.Error("deleting project object failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) ListProjectObjects(ctx context.Context, id common.ByProject, target interface{}) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ?",
				id.GetProjectId(),
			).Find(target)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		s.logger.Error("listing project objects failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) ListAllObjects(ctx context.Context, target interface{}) error {
	err := s.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Find(target)
			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		s.logger.Error("listing all objects failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) DeleteGraphObject(ctx context.Context, id *graph.IDGraph, entity graph.Object) error {
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

type Service struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewService(
	db *gorm.DB,
	logger *zap.SugaredLogger,
) (*Service, error) {
	svc := &Service{
		db, logger,
	}

	svc.logger = svc.logger.With(zap.String("service", "repository"))

	return svc, nil
}
