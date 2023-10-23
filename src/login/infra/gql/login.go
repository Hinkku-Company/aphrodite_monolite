package gql

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/login/usecase"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/graphql/generated"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/schema"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
)

type LoginGQL struct {
	uc usecase.LoginUseCase
}

func NewLoginGQL(uc usecase.LoginUseCase) *LoginGQL {
	return &LoginGQL{uc: uc}
}

func (l *LoginGQL) Login(ctx context.Context, input schema.Credentials) (*schema.Access, error) {
	resp, err := l.uc.GenerateToken(ctx, tables.Credentials{
		Email:    input.UserName,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &schema.Access{
		Token:        resp.Token,
		TokenRefresh: resp.TokenRefresh,
	}, nil
}

func (l *LoginGQL) LogOut(ctx context.Context, input schema.AccessToken) (bool, error) {
	return l.uc.LogOut(ctx, tables.Access{
		Token:        input.Token,
		TokenRefresh: input.TokenRefresh,
	})
}

var _ generated.QueryResolver = (*LoginGQL)(nil)
