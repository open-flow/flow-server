package endpoint

import (
	"autoflow/internal/modules/storage/repo"
	"autoflow/pkg/common"
	"autoflow/pkg/storage/endpoint"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Controller) Save(ctx context.Context, req *endpoint.DBEndpoint) (*endpoint.DBEndpoint, error) {
	entity := &endpoint.DBEndpoint{}
	err := r.repo.SaveProjectObject(ctx, req, entity)
	if err != nil {
		return nil, err
	}
	r.cache.ScheduleRefresh(req)
	return entity, nil
}

func (r *Controller) Delete(ctx context.Context, req *common.IDProject) error {
	entity := &endpoint.DBEndpoint{}
	err := r.repo.DeleteProjectObject(ctx, req, entity)
	if err != nil {
		return err
	}
	r.cache.ScheduleRefresh(req)
	return nil
}

func (r *Controller) List(_ context.Context, req common.ByProject) (*endpoint.Container, error) {
	return r.cache.Get(req)
}

type Controller struct {
	repo   *repo.Service
	logger *zap.Logger
	db     *gorm.DB
	cache  *Cache
}

func NewController(
	repo *repo.Service,
	logger *zap.Logger,
	db *gorm.DB,
	cache *Cache,
) (*Controller, error) {
	svc := &Controller{
		repo, logger, db, cache,
	}
	return svc, nil
}
