package container

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/container/server"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/Skliar-Il/test-task-wallet/pkg/redis"
	"log"
	"os"
	"os/signal"
	"strconv"
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

	kafka, err := logger.NewSyncProducer(&cfg.Logger.KafkaCfg)
	if err != nil {
		log.Fatalf("kafka connection error: %v", err)
	}

	repositoryList := initializer.NewRepositoryList()
	serviceList := initializer.NewServiceList(repositoryList, dbPool)
	app := server.NewServer(cfg, serviceList, redisConn, kafka)

	serverPortStr := strconv.Itoa(int(cfg.Server.PortHttp))
	go func() {
		pid := os.Getpid()
		log.Printf("[PID %d] starting server...", pid)

		if err = app.Listen(":" + serverPortStr); err != nil {
			log.Printf("[PID %d] server listen error: %v", pid, err)
		}
	}()
	select {
	case <-ctx.Done():
		if err = app.Shutdown(); err != nil {
			log.Printf("server shotdown error: %v\n", err)
		}
		if err = logger.SyncMiddleware(); err != nil {
			log.Printf("sync logger error: %v\n", err)
		}
		log.Println("server stopped")

	}
}
