package env

import (
	"fmt"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
)

const (
	portGRPC = "GRPC_PORT"
	hostGRPC = "GRPC_HOST"
)

type GRPCCfgSearcher struct{}

func NewGRPCCfgSearcher() *GRPCCfgSearcher {
	return &GRPCCfgSearcher{}
}

func (s *GRPCCfgSearcher) Get() (*config.GRPCConfig, error) {
	host := os.Getenv(hostGRPC)
	if host == "" {
		return nil, fmt.Errorf("gRPC Host not found")
	}

	port := os.Getenv(portGRPC)
	if port == "" {
		return nil, fmt.Errorf("gRPC Port not found")
	}

	return &config.GRPCConfig{
		Host: host,
		Port: port,
	}, nil
}
