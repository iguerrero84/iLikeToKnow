package initializers

import (
	"context"
	"iLikeToKnow.com/internal/config"
)

type DomainConfig interface {
	config.BaseConfig
	config.PostgresConfig
}

type DomainConfigImpl struct {
	config.BaseConfigImpl
	config.PostgresConfigImpl
}

func LoadConfigDomainService(ctx context.Context) (*DomainConfigImpl, error) {

	cfg, err := config.LoadConfig[DomainConfigImpl](ctx)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
