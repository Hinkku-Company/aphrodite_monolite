package redis

import (
	"context"
	"net"
	"strconv"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/redis/go-redis/v9"
)

type redisConn struct {
	client *redis.Client
	config config.Config
	ctx    context.Context
}

func NewClient(ctx context.Context, config config.Config) *redisConn {
	return &redisConn{
		ctx:    ctx,
		config: config,
	}
}

func (r *redisConn) ConnectRedis() (*redis.Client, error) {
	host := net.JoinHostPort(r.config.RedisHost, r.config.RedisPort)
	logger.Log().Info("Get Redis connection", "url", host)
	db, _ := strconv.Atoi(r.config.RedisDB)
	r.client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: r.config.RedisPassword,
		DB:       db,
	})

	if _, err := r.client.Ping(r.ctx).Result(); err != nil {
		return nil, err
	}

	return r.client, nil
}
