package endpoint

import (
	"autoflow/internal/services/repo"
	"autoflow/pkg/entities/common"
	"autoflow/pkg/entities/endpoint"
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func (c *Cache) ScheduleRefresh(id common.ByProject) {
	c.logger.Info("scheduling sync")
	go func() {
		_ = c.Refresh(context.Background(), id)
	}()
}

func (c *Cache) Refresh(ctx context.Context, id common.ByProject) error {
	logger := c.logger.With(zap.Uint("projectId", id.GetProjectId()))
	lock, err := c.locker.Obtain(ctx, lockKey(id), time.Millisecond*50, &redislock.Options{})
	if err != nil {
		logger.Error("failed to obtain lock", zap.Error(err))
		return err
	}
	defer lock.Release(ctx)

	var endpoints []*endpoint.DBEndpoint
	err = c.repo.ListProjectObjects(ctx, id, &endpoints)
	if err != nil {
		logger.Error("failed to list endpoints", zap.Error(err))
		return err
	}

	err = c.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key(id),
		Value: endpoint.NewContainer(endpoints),
		TTL:   time.Hour * 2,
	})

	if err != nil {
		logger.Error("failed to store cache", zap.Error(err))
		return err
	}

	c.logger.Info("endpoints synchronized")
	return nil
}

func (c *Cache) Get(id common.ByProject) (*endpoint.Container, error) {
	container := &endpoint.Container{}
	err := c.cache.Get(context.Background(), key(id), container)
	if err == cache.ErrCacheMiss {
		err = c.Refresh(context.Background(), id)
		if err != nil {
			return nil, err
		}
		err = c.cache.Get(context.Background(), key(id), container)
	}
	if err != nil {
		c.logger.Error("failed to access cache", zap.Uint("projectId", id.GetProjectId()))
		return nil, err
	}
	return container, nil
}

type Cache struct {
	logger *zap.SugaredLogger
	cache  *cache.Cache
	redis  *redis.Ring
	locker *redislock.Client
	repo   *repo.Service
}

func NewCache(
	logger *zap.SugaredLogger,
	cache *cache.Cache,
	redis *redis.Ring,
	locker *redislock.Client,
	repo *repo.Service,
) (*Cache, error) {
	obj := &Cache{
		logger, cache, redis, locker, repo,
	}
	obj.logger = obj.logger.With(zap.String("cache", "endpoint"))
	return obj, nil
}

func key(id common.ByProject) string {
	return fmt.Sprintf("endpoint.project.%d", id.GetProjectId())
}
func lockKey(id common.ByProject) string {
	return fmt.Sprintf("endpoint.project.%d.lock", id.GetProjectId())
}
