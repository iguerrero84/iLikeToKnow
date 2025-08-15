package config

type BaseConfig interface {
	Local() bool
}

type BaseConfigImpl struct {
	Local_ string `env:"LOCAL"`
}

func (c *BaseConfigImpl) Local() bool {
	return c.Local_ != ""
}
