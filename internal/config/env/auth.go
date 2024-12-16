package env

import (
	"errors"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
	"strconv"
)

const (
	portAuth = "AUTH_PORT"
	hostAuth = "AUTH_HOST"
)

type AuthConfigSearcher struct{}

func NewAuthConfigSearcher() *AuthConfigSearcher {
	return &AuthConfigSearcher{}
}

func (cfg *AuthConfigSearcher) Get() (*config.AuthConfig, error) {
	host := os.Getenv(hostAuth)
	if len(host) == 0 {
		return nil, errors.New("auth host not found")
	}

	portString := os.Getenv(portAuth)
	if len(portString) == 0 {
		return nil, errors.New("auth port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("use integer port")
	}

	return &config.AuthConfig{
		Host: host,
		Port: portString,
	}, nil
}
