package client

import (
	"context"
	"fmt"

	pb "github.com/agusheryanto182/job-recommendations/backend/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client pb.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthClient(address string) (*AuthClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %v", err)
	}

	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) error {
	resp, err := c.client.ValidateRequest(ctx, &pb.ValidateRequestRequest{
		Token: token,
		Guard: "user",
	})
	if err != nil {
		return fmt.Errorf("failed to validate token: %v", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("token validation failed: %s", resp.Error)
	}
	return nil
}

func (c *AuthClient) GetUserID(ctx context.Context, token string) (string, error) {
	resp, err := c.client.GetAuthID(ctx, &pb.GetAuthIDRequest{
		Token: token,
		Guard: "user",
	})
	if err != nil {
		return "", fmt.Errorf("failed to get user ID: %v", err)
	}
	if resp.Error != "" {
		return "", fmt.Errorf("get user ID failed: %s", resp.Error)
	}
	return resp.AuthId, nil
}
