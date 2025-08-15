package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	cnx "iLikeToKnow.com/internal/database/connexion"
	db "iLikeToKnow.com/internal/database/db"
	"log/slog"
)

var (
	logger      = slog.Logger{}
	errNotFound = errors.New("not found")
)

type GetSQLCQueriesObjectFunc[Q db.Queries] func(ctx context.Context, conn *Connection) *Q

type Impl[Q db.Queries] struct {
	connectionPool *pgxpool.Pool
	getQueries     GetSQLCQueriesObjectFunc[Q]
}

func (s Impl[Q]) Close() {
	if s.connectionPool != nil {
		s.connectionPool.Close()
	}
}

func (s *Impl[Q]) TX(ctx context.Context, f QueryFunc[Q]) error {

	return s.Raw(ctx, func(ctx context.Context, conn *Connection, queries *Q) error {
		// begin the transaction
		tx, err := conn.Begin(ctx)
		if err != nil {
			safeErr := fmt.Errorf("failed to begin tx")
			logger.Debug(safeErr.Error(), err)
			return safeErr
		}

		// execute the inner query
		err = f(ctx, conn, queries)
		if err != nil {
			// rollback the tx
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				// only log the error since there's nothing to handle here
				logger.Debug("failed to rollback tx", rollbackErr)
			}

			// return the query error
			safeErr := errors.New("failed to execute tx")
			logger.Debug(safeErr.Error(), err)
			return safeErr
		}

		// commit the transaction, since there's no errors
		err = tx.Commit(ctx)
		if err != nil {
			safeErr := fmt.Errorf("failed to commit tx")
			logger.Debug(safeErr.Error(), err)
			return safeErr
		}

		return nil
	})
}

func (s *Impl[Q]) Raw(ctx context.Context, f QueryFunc[Q]) error {

	// acquire a connection from the pool
	conn, err := s.connectionPool.Acquire(ctx)
	if err != nil {
		safeErr := fmt.Errorf("failed to acquire connection from pool")
		logger.Debug(safeErr.Error(), err)
		return safeErr
	}
	defer conn.Release()

	// get the queries object for SQLC
	queries := s.getQueries(ctx, conn)

	// execute the query
	err = f(ctx, conn, queries)
	if errors.Is(err, pgx.ErrNoRows) {
		// return our own error type to encapsulate the driver-level error
		err = errors.Join(errNotFound)
	}
	if err != nil {
		safeErr := errors.New("failed to execute query")
		logger.Debug(safeErr.Error(), err)
		return err
	}

	return nil
}

// NewDatabase creates a new database interface for easy repository layer interaction
func NewDatabase(
	ctx context.Context,
	dbConfig Config,
) (*Impl[db.Queries], error) {
	return newDatabaseImpl(
		ctx,
		dbConfig,
		func(ctx context.Context, conn *Connection) *db.Queries {
			return db.New(conn)
		},
	)
}

// newDatabaseImpl creates a generically typed database interface
//
// this helper function is a required layer of indirection since the specified getQueries function
// must be declared with a concrete type parameter
func newDatabaseImpl[Q db.Queries](
	ctx context.Context,
	dbConfig Config,
	getQueries GetSQLCQueriesObjectFunc[Q],
) (*Impl[Q], error) {
	// create a new connection pool
	connectionParams := chooseConfig(dbConfig)
	connectionPool, newPoolErr := cnx.NewConnectionPool(
		ctx,
		connectionParams,
	)
	if newPoolErr != nil {
		return nil, errors.New("failed to create connection pool against database")
	}

	// Test the connection
	if err := connectionPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	fmt.Println("successfully connected to database!")

	return &Impl[Q]{
		connectionPool: connectionPool,
		getQueries:     getQueries,
	}, nil
}

func chooseConfig(dbConfig Config) cnx.ConnectionParams {
	var connectionParams cnx.ConnectionParams

	connectionParams = cnx.LocalConnection{
		Username:     dbConfig.PostgresDBUser(),
		Password:     dbConfig.PostgresDBPassword(),
		Host:         dbConfig.PostgresDBHost(),
		Port:         dbConfig.PostgresDBPort(),
		DatabaseName: dbConfig.PostgresDBName(),
	}
	return connectionParams
}

//	type Repo struct {
//		db *sql.DB
//		q  *db.Queries
//		// optional: default per-query timeout
//		timeout time.Duration
//	}
//
// // CreateEvent inserts and returns the created row.
//
//	type CreateEventParams struct {
//		Title       string
//		Description *string
//		StartTime   time.Time
//		EndTime     time.Time
//	}
func (r *Impl[Q]) CreateEvent(ctx context.Context, p db.CreateEventParams) (db.Event, error) {
	id := uuid.New()

	ev, err := r.CreateEvent(ctx, db.CreateEventParams{
		ID:          id,
		Title:       p.Title,
		Description: p.Description,
		//StartTime:   TimeToTimestamptz(p.StartTime),
		//EndTime:     TimeToTimestamptz(p.EndTime),
	})
	if err != nil {
		return db.Event{}, fmt.Errorf("create event: %w", err)
	}
	return ev, nil
}

//
//func (r *Impl[Q]) GetEvent(ctx context.Context, id uuid.UUID) (db.Event, error) {
//
//	ev, err := r.getQueries..GetEvent(ctx, id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return db.Event{}, err
//		}
//		return db.Event{}, fmt.Errorf("get event: %w", err)
//	}
//	return ev, nil
//}
//
//func (r *Impl[Q]) ListEvents(ctx context.Context) ([]db.Event, error) {
//
//	rows, err := r.q.ListEvents(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("list events: %w", err)
//	}
//	return rows, nil
//}
