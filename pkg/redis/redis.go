package redis

import (
	"github.com/gofiber/storage/redis/v3"
)

type Config struct {
	Host string `env:"REDIS_HOST" env-default:"localhost"`
	Port uint16 `env:"REDIS_PORT" env-default:"6379"`
}

func New(cfg Config) (*redis.Storage, error) {
	client := redis.New(redis.Config{
		Host:     cfg.Host,
		Port:     int(cfg.Port),
		Username: "",
		Password: "",
		Database: 0,
	})
	return client, nil
}
