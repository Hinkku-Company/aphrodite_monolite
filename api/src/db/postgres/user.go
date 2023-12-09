package postgres

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
	"github.com/Hinkku-Company/aphrodite_monolite/src/user/repository"
	"github.com/google/uuid"
)

func (q *query) ListUsers(ctx context.Context) ([]tables.User, error) {
	user := new([]tables.User)
	err := q.db.NewSelect().
		Model(user).
		Relation("TypeUser").
		Relation("Credentials").
		Scan(ctx)
	if err != nil {
		q.log.Error("ListUsers", "msg", err.Error())
		return *user, err
	}
	return *user, nil
}

func (q *query) GetUser(ctx context.Context, id uuid.UUID) (tables.User, error) {
	user := new(tables.User)
	err := q.db.NewSelect().
		Model(user).
		Relation("TypeUser").
		Relation("Credentials").
		Where("\"user\".id = ?", id).
		Scan(ctx)
	if err != nil {
		q.log.Error("GetUser", "msg", err.Error())
		return *user, err
	}
	return *user, nil
}

var _ repository.UserRepository = (*query)(nil)
