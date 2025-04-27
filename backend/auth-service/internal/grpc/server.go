package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/agusheryanto182/job-recommendations/backend/proto/auth"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port   int
}

func recoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}

func NewServer(port int) *Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(recoveryInterceptor),
	}
	grpcServer := grpc.NewServer(opts...)
	authServer := NewAuthServer()
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	return &Server{
		server: grpcServer,
		port:   port,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("0.0.0.0:%d", s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("gRPC server starting on port %d", s.port)
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
