package initializer

import (
	"github.com/Skliar-Il/test-task-wallet/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceList struct {
	WalletService service.WalletServiceInterface
}

func NewServiceList(repositories *RepositoryList, dbPool *pgxpool.Pool) *ServiceList {
	return &ServiceList{
		WalletService: service.NewWalletService(dbPool, repositories.WalletRepository),
	}
}
