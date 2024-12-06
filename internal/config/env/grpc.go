package env

import (
	"errors"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
)

const (
	portGRPC = "GRPC_PORT"
	hostGRPC = "GRPC_HOST"
)

type GRPCConfigSearcher struct{}

func NewGRPCConfigSearcher() *GRPCConfigSearcher {
	return &GRPCConfigSearcher{}
}

func (s *GRPCConfigSearcher) Get() (*config.GRPCConfig, error) {
	host := os.Getenv(hostGRPC)
	if host == "" {
		return nil, errors.New("gRPC Host not found")
	}

	port := os.Getenv(portGRPC)
	if port == "" {
		return nil, errors.New("gRPC Port not found")
	}

	return &config.GRPCConfig{
		Host: host,
		Port: port,
	}, nil
}
