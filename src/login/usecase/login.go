package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/auth"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/repository"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
)

type LoginUseCase interface {
	GenerateToken(ctx context.Context, model tables.Credentials) (tables.Access, error)
	LogOut(ctx context.Context, model tables.Access) (bool, error)
}

type loginUseCase struct {
	repoSQL   repository.SQLRepository
	repoRedis repository.RedisRepository
	config    config.Config
	log       *slog.Logger
	auth      *auth.Auth
}

func NewLoginUseCase(
	repoSQL repository.SQLRepository,
	repoRedis repository.RedisRepository,
	config config.Config) LoginUseCase {
	return &loginUseCase{
		config:    config,
		repoSQL:   repoSQL,
		repoRedis: repoRedis,
		log:       logger.Log(),
		auth:      auth.NewAuth(config),
	}
}

type STATUS string

const (
	ACTIVE  STATUS = "active"
	PENDING STATUS = "pending"
	BLOCK   STATUS = "block"
	DELETE  STATUS = "delete"
)

func (l *loginUseCase) GenerateToken(ctx context.Context, model tables.Credentials) (tables.Access, error) {
	// validate input
	// chirris

	// get user credentials
	user, err := l.repoSQL.GetCredentials(ctx, model)
	if err != nil {
		l.log.Error("GetCredentials", "error", err)
		return tables.Access{}, errors.New("invalid credentials")
	}

	// validate password
	if accept, err := l.auth.ValidPasswordHash(model.Password, user.Credentials.Password); err != nil || !accept {
		l.log.Error("GetCredentials", "error", "password")
		return tables.Access{}, errors.New("invalid credentials")
	}

	// validate status
	if user.Credentials.Status.Name != string(ACTIVE) {
		l.log.Error("GetCredentials", "error", "estatus "+user.Credentials.Status.Name)
		return tables.Access{}, errors.New("invalid credentials")
	}

	// generate token
	access, err := l.auth.CreateToken(user)
	if err != nil {
		l.log.Error("GetCredentials", "error", err)
		return tables.Access{}, errors.New("invalid credentials")
	}

	// save in redis
	err = l.repoRedis.SaveToken(ctx, *access)
	if err != nil {
		l.log.Error("GetCredentials", "error", err)
		return tables.Access{}, errors.New("invalid credentials")
	}

	return *access, nil
}

func (l *loginUseCase) LogOut(ctx context.Context, model tables.Access) (bool, error) {
	// remove token
	err := l.repoRedis.RemoveToken(ctx, model.Token)
	if err != nil {
		l.log.Error("RemoveToken", "error", err)
	}

	// remove refresh_token
	err = l.repoRedis.RemoveToken(ctx, model.TokenRefresh)
	if err != nil {
		l.log.Error("RemoveToken", "error", err)
	}

	return true, nil
}

var _ LoginUseCase = &loginUseCase{}
