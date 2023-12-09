package usecase

import (
	"context"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
	"github.com/Hinkku-Company/aphrodite_monolite/src/user/repository"
	"github.com/google/uuid"
)

type UserUseCase interface {
	ListUsers(ctx context.Context) ([]tables.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (tables.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
	log  *slog.Logger
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
		log:  logger.Log(),
	}
}

func (u *userUseCase) ListUsers(ctx context.Context) ([]tables.User, error) {
	resp, err := u.repo.ListUsers(ctx)
	if err != nil {
		u.log.Error("ListUsers", "msg", err)
		return nil, err
	}
	return resp, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id uuid.UUID) (tables.User, error) {
	resp, err := u.repo.GetUser(ctx, id)
	if err != nil {
		u.log.Error("GetUser", "msg", err)
		return resp, err
	}
	return resp, nil
}
