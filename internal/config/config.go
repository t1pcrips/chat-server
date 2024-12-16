package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"net"
	"os"
)

type PgConfig struct {
	Port     int
	Host     string
	User     string
	Name     string
	Password string
}

type GRPCConfig struct {
	Host string
	Port string
}

type HTTPConfig struct {
	Host string
	Port string
}

type SWAGGERConfig struct {
	Host string
	Port string
}

type LogConfig struct {
	LogLevel      zerolog.Level
	LogTimeFormat string
}

type AuthConfig struct {
	Host string
	Port string
}

func (cfg *PgConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}

func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *SWAGGERConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *AuthConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *LogConfig) SetUp() *zerolog.Logger {
	settingsLoger := zerolog.ConsoleWriter{TimeFormat: cfg.LogTimeFormat, Out: os.Stdout, NoColor: false}
	logger := zerolog.New(settingsLoger).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(cfg.LogLevel)
	zerolog.TimeFieldFormat = cfg.LogTimeFormat
	return &logger
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load .env file %w", err)
	}

	return nil
}
