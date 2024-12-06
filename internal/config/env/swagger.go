package env

import (
	"errors"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
	"strconv"
)

const (
	hostSWAGGER = "SWAGGER_HOST"
	portSWAGGER = "SWAGGER_PORT"
)

type SwaggerConfigSearcher struct{}

func NewSwaggerConfigSearcher() *SwaggerConfigSearcher {
	return &SwaggerConfigSearcher{}
}

func (s *SwaggerConfigSearcher) Get() (*config.SWAGGERConfig, error) {
	host := os.Getenv(hostSWAGGER)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	portString := os.Getenv(portSWAGGER)
	if len(portString) == 0 {
		return nil, errors.New("swagger port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("faild to convert swagger port to integer")
	}

	return &config.SWAGGERConfig{
		Host: host,
		Port: portString,
	}, nil
}
