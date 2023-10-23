package repository

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
)

type SQLRepository interface {
	GetCredentials(ctx context.Context, model tables.Credentials) (tables.User, error)
}

type RedisRepository interface {
	SaveToken(ctx context.Context, model tables.Access) error
	GetToken(ctx context.Context, token string) (bool, error)
	RemoveToken(ctx context.Context, token string) error
}
