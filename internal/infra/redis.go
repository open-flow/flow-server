package infra

import (
	"github.com/bsm/redislock"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"time"
)

type RedisDeps struct {
	fx.In
	Config *FlowConfig
}

type RedisProduct struct {
	fx.Out
	Cache  *cache.Cache
	Redis  *redis.Ring
	Locker *redislock.Client
}

func NewRedis(deps RedisDeps) (RedisProduct, error) {
	out := RedisProduct{}

	out.Redis = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"host": deps.Config.RedisHost + ":" + deps.Config.RedisPort,
		},
		Username: deps.Config.RedisUsername,
		Password: deps.Config.RedisPassword,
	})

	out.Locker = redislock.New(out.Redis)

	out.Cache = cache.New(&cache.Options{
		Redis:      out.Redis,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return out, nil
}
