package configs

import "os"

type PortConfig struct {
	PORT string `env:"APP_PORT"`
}

func GetPortConfig() PortConfig {
	return PortConfig{
		PORT: os.Getenv("APP_PORT"),
	}
}
