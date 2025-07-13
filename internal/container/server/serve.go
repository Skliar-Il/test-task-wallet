package server

import (
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/transport/http"
	"github.com/Skliar-Il/test-task-wallet/pkg/exception"
	pkgvalidator "github.com/Skliar-Il/test-task-wallet/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/storage/redis/v3"
)

func NewServer(cfg *config.Config, serviceList *initializer.ServiceList, redisStg *redis.Storage) *fiber.App {

	serverConfig := fiber.Config{
		StructValidator: pkgvalidator.Validator{Validator: validator.New()},
		ErrorHandler:    exception.Middleware,
	}
	server := fiber.New(serverConfig)

	http.NewController(server, cfg, serviceList, redisStg)
	return server

}
