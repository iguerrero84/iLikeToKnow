package connexion

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LocalConnection struct {
	DatabaseName string
	Host         string
	Port         string
	Username     string
	Password     string
}

var _ ConnectionParams = (*LocalConnection)(nil)

func (l LocalConnection) generateConfig(ctx context.Context) (*pgxpool.Config, error) {
	cfg, err := pgxpool.ParseConfig(generateCnxnString(
		host(l.Host),
		port(l.Port),
		dbName(l.DatabaseName),
		user(l.Username),
		password(l.Password),
		noSsl,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection conifguration: %w", err)
	}
	return cfg, nil
}
