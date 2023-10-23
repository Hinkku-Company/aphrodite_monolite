package redis

import (
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/repository"
	"github.com/redis/go-redis/v9"
)

type query struct {
	config config.Config
	log    *slog.Logger
	client *redis.Client
}

func NewQuery(config config.Config, client *redis.Client) *query {
	return &query{
		config: config,
		log:    logger.Log(),
		client: client,
	}
}

var _ repository.RedisRepository = (*query)(nil)
