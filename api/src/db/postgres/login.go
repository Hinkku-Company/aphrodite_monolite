package postgres

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
)

func (q *query) GetCredentials(ctx context.Context, model tables.Credentials) (tables.User, error) {
	user := new(tables.User)
	err := q.db.NewSelect().
		Model(user).
		Relation("Credentials").
		Relation("Credentials.Status").
		Relation("TypeUser").
		Join("LEFT JOIN hk.users_status AS users_status ON users_status.id = credentials.status_id").
		Where("credentials.email = ?", model.Email).
		Scan(ctx)
	if err != nil {
		q.log.Error("GetCredentials", "error", err)
		return tables.User{}, err
	}

	user.AccessRols, err = q.GetRol(ctx, *user)
	if err != nil {
		q.log.Error("GetCredentialsRol", "error", err)
		return tables.User{}, err
	}
	return *user, nil
}

func (q *query) GetRol(ctx context.Context, model tables.User) ([]tables.UserAccessRol, error) {
	var accessRol []tables.UserAccessRol
	err := q.db.NewSelect().
		Model(&accessRol).
		Relation("AccessRol").
		Where("users_id = ?", model.ID).
		Scan(ctx)
	if err != nil {
		q.log.Error("GetCredentialsWithUser", "error", err)
		return accessRol, err
	}

	return accessRol, nil
}
