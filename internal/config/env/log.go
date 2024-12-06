package env

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/t1pcrips/chat-service/internal/config"
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
		return nil, errors.New("logLevel not found")
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return nil, errors.New("failed to conver level to int")
	}

	timeFormat := os.Getenv(logTimeFormat)
	if timeFormat == "" {
		return nil, errors.New("timeFormat not found")
	}

	return &config.LogConfig{
		LogLevel: zerolog.Level(level),
	}, nil
}
