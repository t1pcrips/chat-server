package service

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type FakeTx interface {
	Begin(ctx context.Context) (pgx.Tx, error)

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects

	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row

	Conn() *pgx.Conn
}
