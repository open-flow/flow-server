package sendpoint

import (
	"autoflow/internal/modules/srepo"
	"autoflow/pkg/common"
	"autoflow/pkg/storage/endpoint"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Endpoint) Save(ctx context.Context, req *endpoint.DBEndpoint) (*endpoint.DBEndpoint, error) {
	entity := &endpoint.DBEndpoint{}
	err := r.repo.SaveProjectObject(ctx, req, entity)
	if err != nil {
		return nil, err
	}
	r.cache.ScheduleRefresh(req)
	return entity, nil
}

func (r *Endpoint) Delete(ctx context.Context, req *common.ProjectModel) error {
	entity := &endpoint.DBEndpoint{}
	err := r.repo.DeleteProjectObject(ctx, req, entity)
	if err != nil {
		return err
	}
	r.cache.ScheduleRefresh(req)
	return nil
}

func (r *Endpoint) List(_ context.Context, req common.SpacedObject) (*endpoint.Container, error) {
	return r.cache.Get(req)
}

type Endpoint struct {
	repo   *srepo.Repo
	logger *zap.Logger
	db     *gorm.DB
	cache  *EndpointCache
}

func NewEndpoint(
	repo *srepo.Repo,
	logger *zap.Logger,
	db *gorm.DB,
	cache *EndpointCache,
) (*Endpoint, error) {
	svc := &Endpoint{
		repo, logger, db, cache,
	}
	return svc, nil
}
