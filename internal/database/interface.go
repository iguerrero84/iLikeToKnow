package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db "iLikeToKnow.com/internal/database/db"
)

type Connection = pgxpool.Conn

type QueryFunc[Q db.Queries] func(ctx context.Context, conn *Connection, queries *Q) error

type Config interface {
	Local() bool
	PostgresDBHost() string
	PostgresDBPort() string
	PostgresDBUser() string
	PostgresDBName() string
	PostgresDBPassword() string
}

type Database[Q db.Queries] interface {
	// Close ensures the database is closed cleanly
	Close()
	// TX starts a transaction, and executes the given function in a transaction context
	//
	// The caller should assume that the received Connection and sqlc.Queries are instantiated already,
	// and not manage their lifecycle. The Connection will be closed automatically upon returning.
	// If the passed function returns an error, the transaction is rolled back. Return nil to commit.
	//
	// pgx.ErrNoRows is mapped to our own driver-independent ErrNoRows type
	TX(ctx context.Context, f QueryFunc[Q]) error
	// Raw starts a connection, and executes the given function *without* a transaction context
	//
	// The caller should assume that the received Connection and sqlc.Queries are instantiated already,
	// and not manage their lifecycle. The Connection will be closed automatically upon returning.
	//
	// pgx.ErrNoRows is mapped to our own driver-independent ErrNoRows type
	Raw(ctx context.Context, f QueryFunc[Q]) error
}
