package integration

import (
	"context"
	"fmt"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/container/server"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	redisPkg "github.com/Skliar-Il/test-task-wallet/pkg/redis"
	"github.com/gofiber/fiber/v3"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	baseURL string
	app     *fiber.App
)

var (
	pgContainer    *postgres.PostgresContainer
	redisContainer *redis.RedisContainer
)

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()

	pgContainer, err = postgres.RunContainer(ctx,
		postgres.WithDatabase("wallet_test"),
		postgres.WithUsername("wallet_test"),
		postgres.WithPassword("wallet_test"),
	)
	if err != nil {
		log.Fatalf("start postgres error: %v", err)
	}

	redisContainer, err = redis.RunContainer(ctx)
	if err != nil {
		log.Fatalf("start redis error: %v", err)
	}

	pgHost, err := pgContainer.Host(ctx)
	if err != nil {
		log.Fatalf("get postgres host error: %v", err)
	}
	pgPort, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("get postgres port error: %v", err)
	}

	redisHost, err := redisContainer.Host(ctx)
	if err != nil {
		log.Fatalf("get redis host error: %v", err)
	}
	redisPort, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		log.Fatalf("get redis port error: %v", err)
	}

	pgPortStr := pgPort.Port()
	pgPortInt, err := strconv.Atoi(pgPortStr)
	if err != nil {
		log.Fatalf("convert postgres port to int error: %v", err)
	}

	redisPortStr := redisPort.Port()
	redisPortInt, err := strconv.Atoi(redisPortStr)
	if err != nil {
		log.Fatalf("convert redis port to int error: %v", err)
	}

	cfg := &config.Config{
		Server: config.Server{
			Version: 1,
		},
		DataBase: database.Config{
			Host:     pgHost,
			Port:     uint16(pgPortInt),
			User:     "wallet_test",
			Password: "wallet_test",
			Name:     "wallet_test",
			MinConn:  1,
			MaxConn:  5,
		},
		Redis: redisPkg.Config{
			Host: redisHost,
			Port: uint16(redisPortInt),
		},
		Logger: logger.Config{Mode: "debug"},
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

LOOP:
	for {
		select {
		case <-timeoutCtx.Done():
			log.Fatalf("database connection timeout error")
		case <-time.After(1 * time.Second):
			if database.Ping(ctx, cfg.DataBase) {
				break LOOP
			}
		}
	}
	for {
		stateRedis, err := redisContainer.State(ctx)
		if err != nil {
			log.Fatalf("check redis container error: %v", err)
		}
		if stateRedis.Running {
			break
		}
		time.Sleep(1 * time.Second)
	}

	db, err := database.New(ctx, cfg.DataBase)
	if err != nil {
		log.Fatalf("connect db error: %v", err)
	}
	rds, err := redisPkg.New(cfg.Redis)
	if err != nil {
		log.Fatalf("connect redis error: %v", err)
	}

	services := initializer.NewServiceList(initializer.NewRepositoryList(), db)
	app = server.NewServer(cfg, services, rds)

	port := 9999
	baseURL = fmt.Sprintf("http://localhost:%d/api/v1", port)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()

	code := m.Run()

	_ = app.Shutdown()
	_ = pgContainer.Terminate(ctx)
	_ = redisContainer.Terminate(ctx)

	os.Exit(code)
}
