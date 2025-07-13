package http

import (
	"fmt"
	"github.com/Flussen/swagger-fiber-v3"
	"github.com/Skliar-Il/test-task-wallet/docs"
	_ "github.com/Skliar-Il/test-task-wallet/docs"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/storage/redis/v3"
	"strconv"
	"time"
)

func NewController(server *fiber.App, cfg *config.Config, services *initializer.ServiceList, redisStg *redis.Storage) {
	server.Use(cors.New())
	server.Use(logger.Middleware(&cfg.Logger))
	server.Use(cache.New(cache.Config{
		Storage:      redisStg,
		Expiration:   10 * time.Second,
		CacheControl: true,
	}))

	api := server.Group(fmt.Sprintf("/api/v%d", cfg.Server.Version))
	api.Use("/swagger/*", swagger.HandlerDefault)
	docs.SwaggerInfo.Version = strconv.Itoa(cfg.Server.Version)
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/api/v%d", cfg.Server.Version)

	walletHandler := NewWalletHandler(services.WalletService)
	api.Post("/wallet", walletHandler.UpdateWallet)
	api.Get("/wallets/:WALLET_UUID", walletHandler.GetWallet)

	api.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.Status(200).SendString("pong")
	})

}
