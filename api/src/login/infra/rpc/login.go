package rpc

import (
	"context"
	"log/slog"

	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/login/usecase"
	pb "github.com/Hinkku-Company/aphrodite_monolite/src/shared/grpc/v1/proto"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Login struct {
	uc     usecase.LoginUseCase
	server *grpc.Server
	log    *slog.Logger
	pb.UnimplementedLoginServiceServer
}

func NewLogin(server *grpc.Server, uc usecase.LoginUseCase) *Login {
	return &Login{
		server: server,
		uc:     uc,
		log:    logger.Log(),
	}
}

func (l *Login) RegisterService() {
	pb.RegisterLoginServiceServer(l.server, l)
}

func (l *Login) GetLogin(ctx context.Context, input *pb.CredentialsInput) (*pb.AccessResponse, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	// }
	// l.log.Info("login", "ctx", md)
	resp, err := l.uc.GenerateToken(ctx, tables.Credentials{
		Email:    input.UserName,
		Password: input.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	return &pb.AccessResponse{
		Token:        resp.Token,
		TokenRefresh: resp.TokenRefresh,
	}, nil
}

func (l *Login) LogOut(ctx context.Context, input *pb.TokenAccessInput) (*pb.Empty, error) {
	_, _ = l.uc.LogOut(ctx, tables.Access{
		Token:        input.Token,
		TokenRefresh: input.TokenRefresh,
	})
	return &pb.Empty{}, nil
}
