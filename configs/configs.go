package configs

import (
	"os"
)

const (
	prod = "production"
)

type Config struct {
	Env     string        `env:"env"`
	MongoDB MongoDBConfig `json:"mongodb"`
	Port    PortConfig    `env:"APP_PORT"`
}

func (c Config) IsProd() bool {
	return c.Env == prod
}

func GetConfig() Config {
	return Config{
		Env:     os.Getenv("ENV"),
		MongoDB: GetMongoDBConfig(),
		Port:    GetPortConfig(),
	}
}
