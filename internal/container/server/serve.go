package server

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/transport/http"
	"github.com/Skliar-Il/test-task-wallet/pkg/exception"
	pkgvalidator "github.com/Skliar-Il/test-task-wallet/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"log"
	"os/signal"
	"syscall"
)

func Serve(cfg *config.Config, serviceList *initializer.ServiceList) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	serverConfig := fiber.Config{
		StructValidator: pkgvalidator.Validator{Validator: validator.New()},
		ErrorHandler:    exception.Middleware,
	}
	listenConfig := fiber.ListenConfig{
		EnablePrefork: true,
	}
	server := fiber.New(serverConfig)

	http.NewController(server, cfg, serviceList)

	go func() {
		if err := server.Listen(":8080", listenConfig); err != nil {
			log.Fatalf("start server error: %v", err)
		}
	}()
	select {
	case <-ctx.Done():
		if err := server.Shutdown(); err != nil {
			log.Fatalf("server shotdown error: %v", err)
		}
		log.Printf("server stoped")
	}
}
