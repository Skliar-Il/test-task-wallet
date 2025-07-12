package container

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/config"
	"github.com/Skliar-Il/test-task-wallet/internal/container/initializer"
	"github.com/Skliar-Il/test-task-wallet/internal/container/server"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"log"
)

func NewApp() {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("get config error: %v", err)
	}

	dbPool, err := database.New(ctx, cfg.DataBase)
	if err != nil {
		log.Fatalf("init database error: %v", err)
	}
	defer dbPool.Close()

	repositoryList := initializer.NewRepositoryList()
	serviceList := initializer.NewServiceList(repositoryList, dbPool)
	server.Serve(cfg, serviceList)
}
