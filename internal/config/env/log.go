package env

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
	"strconv"
)

type LogCfSearcher struct{}

func NewLogCfgSearcher() *LogCfSearcher {
	return &LogCfSearcher{}
}

const (
	logLevel      = "LOG_LEVEL"
	logTimeFormat = "LOG_TIME_FORMAT"
)

func (s *LogCfSearcher) Get() (*config.LogConfig, error) {
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
