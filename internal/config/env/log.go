package env

import (
	"chat-server/internal/config"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strconv"
)

type LogConfigSearcher struct{}

func NewLogConfigSearcher() *LogConfigSearcher {
	return &LogConfigSearcher{}
}

const (
	logLevel      = "LOG_LEVEL"
	logTimeFormat = "LOG_TIME_FORMAT"
)

func (s *LogConfigSearcher) Get() (*config.LogConfig, error) {
	levelStr := os.Getenv(logLevel)
	if levelStr == "" {
		return nil, fmt.Errorf("logLevel not found")
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return nil, fmt.Errorf("failed to conver level to int: %w", err)
	}

	timeFormat := os.Getenv(logTimeFormat)
	if timeFormat == "" {
		return nil, fmt.Errorf("timeFormat not found")
	}

	return &config.LogConfig{
		LogLevel: zerolog.Level(level),
	}, nil
}
