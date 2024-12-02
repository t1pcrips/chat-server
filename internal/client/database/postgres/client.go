package postgres

import (
	"chat-server/internal/client/database"
	"chat-server/pkg/errs"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	masterDBC database.DB
}

func New(ctx context.Context, dsn string) (database.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errs.ErrConnectToDb
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() database.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}
	return nil
}
