package postgres

import (
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/repository"
	"github.com/uptrace/bun"
)

type query struct {
	config config.Config
	db     *bun.DB
	log    *slog.Logger
}

func NewQuery(config config.Config, db *bun.DB) *query {
	return &query{
		config: config,
		db:     db,
		log:    logger.Log(),
	}
}

var _ repository.SQLRepository = (*query)(nil)
