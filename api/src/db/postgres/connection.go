package postgres

import (
	"context"
	"embed"
	"fmt"
	"net"
	"time"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
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

func (p *postgresConn) parseURL() *pgx.ConnConfig {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		p.config.DBUser,
		p.config.DBPassword,
		net.JoinHostPort(p.config.DBHost, p.config.DBPort),
		p.config.DBName)

	if !p.config.DBIsInsecure {
		connStr += "?sslmode=require"
	}

	config, err := pgx.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}
	return config
}

func (p *postgresConn) ConnectPostgres() (*bun.DB, error) {
	host := net.JoinHostPort(p.config.DBHost, p.config.DBPort)
	logger.Log().Info("Get Postgres connection", "url", host)
	sqldb := stdlib.OpenDB(*p.parseURL())
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
