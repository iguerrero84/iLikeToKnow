package initializers

import (
	"context"
	"iLikeToKnow.com/internal/database"
	db "iLikeToKnow.com/internal/database/db"
	"iLikeToKnow.com/internal/domain"
)

var (
	configInitializerFunc = func(ctx context.Context) (DomainConfig, error) {
		return LoadConfigDomainService(ctx)
	}

	databaseServiceInitializerFunc = func(ctx context.Context, dbConfig database.Config,
	) (database.Database[db.Queries], error) {
		return database.NewDatabase(ctx, dbConfig)
	}
)

type Dependencies struct {
	DomainService   domain.Service
	DatabaseService IDatabaseService
}

// NewDefaultDomainService is a reusable method to create a domain service
// It provides the domain service, as well as its dependencies
//
// The dependencies are required since typically health checking a user of the domain service will entail
// validating that it's sub-services are functioning correctly
func NewDefaultDomainService(ctx context.Context) (Dependencies, error) {

	appConfig, err := configInitializerFunc(ctx)
	if err != nil {
		return Dependencies{}, err
	}

	dbService, err := databaseServiceInitializerFunc(ctx, appConfig)
	if err != nil {
		return Dependencies{}, err
	}

	domainService := domain.NewService(
		dbService,
	)

	return Dependencies{
		DomainService: domainService,
	}, nil
}
