package config

type PostgresConfig interface {
	PostgresDBHost() string
	PostgresDBPort() string
	PostgresDBUser() string
	PostgresDBName() string
	PostgresDBPassword() string
}
type PostgresConfigImpl struct {
	PostgresDBHost_ string `env:"POSTGRES_DB_HOST,notEmpty"`
	PostgresDBPort_ string `env:"POSTGRES_DB_PORT,notEmpty"`
	PostgresDBUser_ string `env:"POSTGRES_DB_USER,notEmpty"`
	PostgresDBName_ string `env:"POSTGRES_DB_NAME,notEmpty"`

	// PostgresDBPassword is only used to connect to a local instance for testing.
	PostgresDBPassword_ string `env:"POSTGRES_DB_PASSWORD"`
}

func (c *PostgresConfigImpl) PostgresDBHost() string {
	return c.PostgresDBHost_
}

func (c *PostgresConfigImpl) PostgresDBPort() string {
	return c.PostgresDBPort_
}

func (c *PostgresConfigImpl) PostgresDBUser() string {
	return c.PostgresDBUser_
}

func (c *PostgresConfigImpl) PostgresDBName() string {
	return c.PostgresDBName_
}

func (c *PostgresConfigImpl) PostgresDBPassword() string {
	return c.PostgresDBPassword_
}
