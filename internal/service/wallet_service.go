package service

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/dto"
	"github.com/Skliar-Il/test-task-wallet/internal/repository"
	"github.com/Skliar-Il/test-task-wallet/pkg/database"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/Skliar-Il/test-task-wallet/pkg/render"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type WalletServiceInterface interface {
	GetWallet(ctx context.Context, id uuid.UUID) (*dto.WalletDTO, error)
	UpdateWallet(ctx context.Context, data *dto.UpdateWalletDTO) (*dto.WalletDTO, error)
}

type WalletService struct {
	dbPool           *pgxpool.Pool
	walletRepository repository.WalletRepositoryInterface
}

func NewWalletService(dbPool *pgxpool.Pool, walletRepository repository.WalletRepositoryInterface) *WalletService {
	return &WalletService{
		dbPool:           dbPool,
		walletRepository: walletRepository,
	}
}

func (s *WalletService) GetWallet(ctx context.Context, id uuid.UUID) (*dto.WalletDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetWallet")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	wallet, err := s.walletRepository.GetWalletById(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "wallet not found")
			return nil, render.Error(fiber.ErrNotFound, "wallet not found")
		}

		localLogger.Error(ctx, "get wallet error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}

	localLogger.Info(ctx, "finish srv func GetWallet")
	return wallet, nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, data *dto.UpdateWalletDTO) (*dto.WalletDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func UpdateWallet")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	wallet, err := s.walletRepository.GetWalletById(ctx, tx, data.WalletId)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {

			localLogger.Info(ctx, "wallet not found")
			wallet, err = s.walletRepository.CreateWallet(ctx, tx, data.WalletId)
			if err != nil {
				localLogger.Error(ctx, "create wallet error", zap.Error(err))
				return nil, render.Error(fiber.ErrInternalServerError, "")

			}
			localLogger.Info(ctx, "wallet created")
		} else {

			localLogger.Error(ctx, "get wallet error", zap.Error(err))
			return nil, render.Error(fiber.ErrInternalServerError, "")
		}
	}
	localLogger.Info(ctx, "get wallet")

	if data.OperationType == "WITHDRAW" {
		if wallet.Amount < data.Amount {
			localLogger.Info(ctx, "amount on wallet less then amount WITHDRAW")
			return nil, render.Error(fiber.ErrBadRequest, "insufficient funds")
		}
		data.Amount *= -1
	}

	wallet, err = s.walletRepository.UpdateWallet(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "update wallet error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	localLogger.Info(ctx, "wallet updated")

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
	}

	localLogger.Info(ctx, "finish srv func UpdateWallet")
	return wallet, nil
}
