package grpc

import "cv-service/internal/grpc/client"

type Config struct {
	AuthServiceAddress string
}

type Clients struct {
	Auth *client.AuthClient
}

func NewClients(cfg *Config) (*Clients, error) {
	authClient, err := client.NewAuthClient(cfg.AuthServiceAddress)
	if err != nil {
		return nil, err
	}

	return &Clients{
		Auth: authClient,
	}, nil
}
