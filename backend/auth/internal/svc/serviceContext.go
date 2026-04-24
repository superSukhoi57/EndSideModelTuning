package svc

import (
	"auth/internal/config"
	"auth/internal/dao"

	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config   config.Config
	Rdb      *redis.Client
	StateDAO *dao.StateDAO
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})

	return &ServiceContext{
		Config:   c,
		Rdb:      rdb,
		StateDAO: dao.NewStateDAO(rdb),
	}
}
