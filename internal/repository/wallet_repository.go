package repository

import (
	"context"
	"github.com/Skliar-Il/test-task-wallet/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type WalletRepositoryInterface interface {
	GetWalletById(ctx context.Context, tx pgx.Tx, walletId uuid.UUID) (*dto.WalletDTO, error)
	CreateWallet(ctx context.Context, tx pgx.Tx, walletId uuid.UUID) (*dto.WalletDTO, error)
	UpdateWallet(ctx context.Context, tx pgx.Tx, date *dto.UpdateWalletDTO) (*dto.WalletDTO, error)
}

type WalletRepository struct {
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{}
}

func (WalletRepository) GetWalletById(ctx context.Context, tx pgx.Tx, walletId uuid.UUID) (*dto.WalletDTO, error) {
	query := `SELECT id, amount FROM wallet WHERE id = $1`
	var wallet dto.WalletDTO
	if err := tx.QueryRow(ctx, query, walletId).Scan(&wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (WalletRepository) CreateWallet(ctx context.Context, tx pgx.Tx, walletId uuid.UUID) (*dto.WalletDTO, error) {
	query := `INSERT INTO wallet (id, amount) VALUES ($1, 0) RETURNING id, amount`
	var wallet dto.WalletDTO
	if err := tx.QueryRow(ctx, query, walletId).Scan(&wallet.WalletId, &wallet.Amount); err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (WalletRepository) UpdateWallet(ctx context.Context, tx pgx.Tx, date *dto.UpdateWalletDTO) (*dto.WalletDTO, error) {
	query := `UPDATE wallet SET amount = amount + $1 WHERE id = $2 RETURNING id, amount`
	var wallet dto.WalletDTO
	if err := tx.QueryRow(ctx, query, date.Amount, date.WalletId).Scan(&wallet.WalletId, &wallet.Amount); err != nil {
		return nil, err
	}
	return &wallet, nil
}
