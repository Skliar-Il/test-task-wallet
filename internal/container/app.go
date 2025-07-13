package container

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/container/server"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"github.com/Skliar-Il/test-task-wallet/pkg/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewApp() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("get config error: %v", err)
	}

	dbPool, err := database.New(ctx, cfg.DataBase)
	if err != nil {
		log.Fatalf("init database error: %v", err)
	}
	defer dbPool.Close()

	redisConn, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatalf("redis connection error: %v", err)
	}

	repositoryList := initializer.NewRepositoryList()
	serviceList := initializer.NewServiceList(repositoryList, dbPool)
	app := server.NewServer(cfg, serviceList, redisConn)

	go func() {
		pid := os.Getpid()
		log.Printf("[PID %d] starting server...", pid)

		if err := app.Listen(":8080"); err != nil {
			log.Printf("[PID %d] server listen error: %v", pid, err)
		}
	}()
	select {
	case <-ctx.Done():
		if err := app.Shutdown(); err != nil {
			log.Fatalf("server shotdown error: %v", err)
		}
		log.Printf("server stoped")
	}
}
