package env

import (
	"fmt"
	"github.com/t1pcrips/chat-service/internal/config"
	"os"
	"strconv"
)

const (
	portPg   = "PG_PORT"
	hostPg   = "PG_HOST"
	name     = "POSTGRES_DB"
	user     = "POSTGRES_USER"
	password = "POSTGRES_PASSWORD"
)

type PgCfgSearcher struct{}

func NewPgCfgSearcher() *PgCfgSearcher {
	return &PgCfgSearcher{}
}

func (s *PgCfgSearcher) Get() (*config.PgConfig, error) {
	dbHost := os.Getenv(hostPg)
	if dbHost == "" {
		return nil, fmt.Errorf("dbHost not found")
	}

	dbPortStr := os.Getenv(portPg)
	if dbPortStr == "" {
		return nil, fmt.Errorf("dbPort not found")
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("faild to convert dbPort to integer")
	}

	dbName := os.Getenv(name)
	if dbName == "" {
		return nil, fmt.Errorf("dbName not found")
	}

	dbUser := os.Getenv(user)
	if dbUser == "" {
		return nil, fmt.Errorf("dbUser not found")
	}

	dbPassword := os.Getenv(password)
	if dbPassword == "" {
		return nil, fmt.Errorf("dbPassword not found")
	}

	return &config.PgConfig{
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
	}, nil
}
