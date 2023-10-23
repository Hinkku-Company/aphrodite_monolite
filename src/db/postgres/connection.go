package postgres

import (
	"context"
	"database/sql"
	"embed"
	"net"
	"time"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type postgresConn struct {
	client *bun.DB
	config config.Config
	ctx    context.Context
}

func NewClient(ctx context.Context, config config.Config) *postgresConn {
	return &postgresConn{
		ctx:    ctx,
		config: config,
	}
}

func (p *postgresConn) ConnectPostgres() (*bun.DB, error) {
	host := net.JoinHostPort(p.config.DBHost, p.config.DBPort)
	logger.Log().Info("Get Postgres connection", "url", host)
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(host),
		pgdriver.WithUser(p.config.DBUser),
		pgdriver.WithPassword(p.config.DBPassword),
		pgdriver.WithDatabase(p.config.DBName),
		pgdriver.WithApplicationName("Aphrodite"),
		pgdriver.WithInsecure(p.config.DBIsInsecure),
	))
	p.client = bun.NewDB(sqldb, pgdialect.New())

	p.client.SetMaxOpenConns(25)
	p.client.SetMaxIdleConns(25)
	p.client.SetConnMaxLifetime(5 * time.Minute)

	if p.config.AppENV == "development" {
		p.client.AddQueryHook(bundebug.NewQueryHook())
		bundebug.NewQueryHook(bundebug.WithVerbose(true))
	}

	if err := p.client.Ping(); err != nil {
		return nil, err
	}

	return p.client, nil
}

func (p *postgresConn) MigrationsUP() error {

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(p.client.DB, "migrations"); err != nil {
		return err
	}

	return nil
}
