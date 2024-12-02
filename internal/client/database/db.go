package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Handler func(ctx context.Context) error

type Client interface {
	DB() DB
	Close() error
}

type TxManeger interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Query struct {
	Name     string
	QueryRow string
}

type SqlExecer interface {
	NamedExecer
	QueryExecer
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, deps interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, deps interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SqlExecer
	Transactor
	Pinger
	Close()
}
