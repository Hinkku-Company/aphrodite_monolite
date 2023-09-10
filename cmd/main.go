package main

import (
	"context"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/db/postgres"
	"github.com/Hinkku-Company/aphrodite_monolite/src/db/redis"
)

func main() {
	conf, err := config.NewConfig().LoadConfigFromEnv()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// redis
	_, err = redis.NewClient(context.TODO(), conf).ConnectRedis()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// postgres
	_, err = postgres.NewClient(context.Background(), conf).ConnectPostgres()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	logger.Log().Info("", "msg", conf.AppENV)
}
