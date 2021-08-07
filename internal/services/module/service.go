package module

import (
	"autoflow/pkg/entities/module"
	"autoflow/pkg/topics"
	"context"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

type Service struct {
	db     *gorm.DB
	logger *zap.Logger
	nc     *nats.Conn
}

func New(db *gorm.DB, logger *zap.Logger, nc *nats.Conn) *Service {
	svc := &Service{
		db:     db,
		nc:     nc,
		logger: logger.With(zap.String("service", "modules")),
	}
	return svc
}

func (s *Service) Save(ctx context.Context, req *module.DBModule) (*module.DBModule, error) {
	entity := &module.DBModule{
		ID: req.ID,
	}

	err := s.db.
		Session(&gorm.Session{Context: ctx}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Assign(req).FirstOrCreate(entity)
			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	s.notifyChange()

	return entity, nil
}

func (s *Service) Delete(ctx context.Context, req *module.DeleteRequest) (*module.DeleteResponse, error) {
	err := s.db.
		Session(&gorm.Session{Context: ctx}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Delete(&module.DBModule{ID: req.ID})
			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	s.notifyChange()

	return &module.DeleteResponse{}, nil
}

func (s *Service) List(ctx context.Context, _ *module.ListRequest) (*module.ListResponse, error) {
	var list []module.DBModule

	err := s.db.
		Session(&gorm.Session{Context: ctx}).
		Transaction(func(tx *gorm.DB) error {
			log.Default()
			res := tx.Model(&module.DBModule{}).Find(&list)
			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &module.ListResponse{
		Modules: list,
	}, nil
}

func (s *Service) notifyChange() {
	err := s.nc.Publish(topics.MODULES_SYNC, make([]byte, 0))
	if err != nil {
		s.logger.Error("publish failed", zap.Error(err))
	}
}
