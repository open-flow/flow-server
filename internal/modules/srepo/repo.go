package srepo

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Repo) ListAllObjects(ctx context.Context, target interface{}) error {
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

type Repo struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewRepo(
	db *gorm.DB,
	logger *zap.SugaredLogger,
) (*Repo, error) {
	svc := &Repo{
		db, logger,
	}

	svc.logger = svc.logger.With(zap.String("service", "repository"))

	return svc, nil
}
