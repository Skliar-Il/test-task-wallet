package dto

import "github.com/google/uuid"

type UpdateWalletDTO struct {
	WalletId      uuid.UUID `json:"valletId" validate:"uuid4" format:"uuid"`
	OperationType string    `json:"operationType" validate:"oneof=DEPOSIT WITHDRAW"`
	Amount        int       `json:"amount" validate:"gte=0"`
}

type WalletDTO struct {
	WalletId uuid.UUID `json:"walletId" validate:"uuid4" format:"uuid"`
	Amount   int       `json:"amount" validate:"gte=0"`
}
