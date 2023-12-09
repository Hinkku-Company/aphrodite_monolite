package repository

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
	"github.com/google/uuid"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]tables.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (tables.User, error)
}
