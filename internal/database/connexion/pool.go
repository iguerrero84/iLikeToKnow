package connexion

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionParams interface {
	generateConfig(ctx context.Context) (*pgxpool.Config, error)
}

type ConnectionPoolOptions interface {
	placeholder()
}

func NewConnectionPool(
	ctx context.Context,
	connection ConnectionParams,
) (*pgxpool.Pool, error) {
	cfg, err := connection.generateConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("error generating connection components: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(
		ctx,
		cfg,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}
	return pool, nil
}
