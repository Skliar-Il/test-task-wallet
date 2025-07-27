package server

import (
	"github.com/IBM/sarama"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/transport/http"
	"github.com/Skliar-Il/test-task-wallet/pkg/exception"
	pkgvalidator "github.com/Skliar-Il/test-task-wallet/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/storage/redis/v3"
	"time"
)

func NewServer(cfg *config.Config, serviceList *initializer.ServiceList, redisStg *redis.Storage, broker sarama.SyncProducer) *fiber.App {

	serverConfig := fiber.Config{
		AppName:         "Wallet",
		StructValidator: pkgvalidator.Validator{Validator: validator.New()},
		ErrorHandler:    exception.Middleware,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     30 * time.Second,
		ProxyHeader:     fiber.HeaderXForwardedFor,
	}
	server := fiber.New(serverConfig)

	http.NewController(server, cfg, serviceList, redisStg, broker)
	return server

}
