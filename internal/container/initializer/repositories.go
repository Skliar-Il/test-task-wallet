package initializer

import "github.com/Skliar-Il/test-task-wallet/internal/repository"

type RepositoryList struct {
	WalletRepository repository.WalletRepositoryInterface
}

func NewRepositoryList() *RepositoryList {
	return &RepositoryList{
		WalletRepository: repository.NewWalletRepository(),
	}
}
