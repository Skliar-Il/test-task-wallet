package config

import (
	"fmt"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/Skliar-Il/test-task-wallet/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Server struct {
	Version int `env:"SERVER_VERSION"`
}

type Config struct {
	Server   Server          `env:"SERVER"`
	DataBase database.Config `env:"POSTGRES"`
	Redis    redis.Config    `env:"REDIS"`
	Logger   logger.Config   `env:"LOGGER"`
}

func New() (*Config, error) {
	var cfg Config

	if err := godotenv.Load("config.env"); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("load config.env error: %w", err)
		}

		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("load .env error: %w", err)
		}
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
