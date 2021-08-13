package srepo

import (
	"autoflow/pkg/common"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Repo) SaveProjectObject(ctx context.Context, save common.ProjectObject, entity common.ProjectObject) error {
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

func (s *Repo) GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
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

func (s *Repo) DeleteProjectObject(ctx context.Context, id *common.IDProject, entity common.ProjectObject) error {
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

func (s *Repo) ListProjectObjects(ctx context.Context, id common.ByProject, target interface{}) error {
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
