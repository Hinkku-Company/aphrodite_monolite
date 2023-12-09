package rpc

import (
	"context"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	pb "github.com/Hinkku-Company/aphrodite_monolite/src/shared/grpc/v1/proto"
	"github.com/Hinkku-Company/aphrodite_monolite/src/user/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	uc     usecase.UserUseCase
	log    *slog.Logger
	server *grpc.Server
	pb.UnimplementedUserServicesServer
}

func NewUser(server *grpc.Server, uc usecase.UserUseCase) *User {
	return &User{
		uc:     uc,
		server: server,
		log:    logger.Log(),
	}
}

func (u *User) RegisterService() {
	pb.RegisterUserServicesServer(u.server, u)
}

func (u *User) GetUser(ctx context.Context, input *pb.UserInput) (*pb.UserResponse, error) {
	uid, err := uuid.Parse(input.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "uuid for user not valid")
	}
	result, err := u.uc.GetUser(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.UserResponse{
		Id:       result.ID.String(),
		Name:     result.Name,
		TypeUser: result.TypeUser.Name,
		Email:    result.Credentials.Email,
	}, nil
}

func (u *User) ListUser(ctx context.Context, input *pb.Empty) (*pb.ListUserResponse, error) {
	result, err := u.uc.ListUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var list []*pb.UserResponse
	for _, user := range result {
		list = append(list, &pb.UserResponse{
			Id:       user.ID.String(),
			Name:     user.Name,
			TypeUser: user.TypeUser.Name,
			Email:    user.Credentials.Email,
		})
	}
	return &pb.ListUserResponse{
		Items: list,
	}, nil
}
