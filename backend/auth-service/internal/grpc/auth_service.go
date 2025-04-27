package grpc

import (
	"auth-service/internal/auth"
	"context"

	pb "github.com/agusheryanto182/job-recommendations/backend/proto/auth"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) ValidateRequest(ctx context.Context, req *pb.ValidateRequestRequest) (*pb.ValidateRequestResponse, error) {
	_, err := auth.ParseJwt(req.Token, auth.User)
	if err != nil {
		return &pb.ValidateRequestResponse{
			Token: "",
			Error: err.Error(),
		}, nil
	}

	err = auth.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateRequestResponse{
			Token: "",
			Error: err.Error(),
		}, nil
	}

	return &pb.ValidateRequestResponse{
		Token: req.Token,
		Error: "",
	}, nil
}

func (s *AuthServer) GetAuthID(ctx context.Context, req *pb.GetAuthIDRequest) (*pb.GetAuthIDResponse, error) {
	claim, err := auth.ParseJwt(req.Token, auth.Guard(req.Guard))
	if err != nil {
		return &pb.GetAuthIDResponse{
			AuthId: "",
			Error:  err.Error(),
		}, nil
	}

	return &pb.GetAuthIDResponse{
		AuthId: claim.Id,
		Error:  "",
	}, nil
}
