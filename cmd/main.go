package main

import (
	"context"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/db/postgres"
	"github.com/Hinkku-Company/aphrodite_monolite/src/db/redis"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/infra/rpc"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/usecase"
)

func main() {
	conf, err := config.NewConfig().LoadConfigFromEnv()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	logger.Log().Info("Environment", "msg", conf.AppENV)

	// redis
	rr, err := redis.NewClient(context.TODO(), conf).ConnectRedis()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// postgres
	psqlConn := postgres.NewClient(context.Background(), conf)
	psql, err := psqlConn.ConnectPostgres()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if err := psqlConn.MigrationsUP(); err != nil {
		slog.Error(err.Error())
		return
	}

	// repository
	psqlRepo := postgres.NewQuery(conf, psql)
	rrRepo := redis.NewQuery(conf, rr)

	// DI
	// login
	loginUC := usecase.NewLoginUseCase(psqlRepo, rrRepo, conf)

	// server
	server := NewAPIServer(conf)
	rpc.NewLogin(server.grpcServer, loginUC).RegisterService()
	// run
	server.Run()
}
