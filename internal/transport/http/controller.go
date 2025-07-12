package http

import (
	"fmt"
	swagger "github.com/Flussen/swagger-fiber-v3"
	_ "github.com/Skliar-Il/test-task-wallet/docs"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewController(server *fiber.App, cfg *config.Config, services *initializer.ServiceList) {
	server.Use(cors.New())
	server.Use(logger.Middleware(&cfg.Logger))

	api := server.Group(fmt.Sprintf("/api/v%d", cfg.Server.Version))
	api.Use("/swagger/*", swagger.HandlerDefault)

	walletHandler := NewWalletHandler(services.WalletService)
	api.Post("/wallet", walletHandler.UpdateWallet)
	api.Get("/wallets/:WALLET_UUID", walletHandler.GetWallet)

	api.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.Status(200).SendString("pong")
	})
}
