package main

import (
	"log/slog"
	"net"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	config     config.Config
	lis        net.Listener
	log        *slog.Logger
	grpcServer *grpc.Server
}

func NewAPIServer(config config.Config) *server {
	return &server{
		config:     config,
		log:        logger.Log(),
		grpcServer: grpc.NewServer(),
	}
}

func (s *server) Run() {
	url := net.JoinHostPort("0.0.0.0", s.config.APPPort)
	logger.Log().Info("Starting server", "host", "http://"+url, "mode", s.config.AppENV)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		s.log.Error("Failed to listen", "MSG", err)
		return
	}
	s.lis = lis

	reflection.Register(s.grpcServer)
	if err := s.grpcServer.Serve(lis); err != nil {
		s.log.Error("Failed to start server", "MSG", err)
		return
	}
}
