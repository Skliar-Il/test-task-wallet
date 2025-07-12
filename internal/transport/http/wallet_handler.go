package http

import (
	"github.com/Skliar-Il/test-task-wallet/internal/dto"
	"github.com/Skliar-Il/test-task-wallet/internal/service"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/Skliar-Il/test-task-wallet/pkg/render"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type WalletHandler struct {
	walletService service.WalletServiceInterface
}

func NewWalletHandler(walletService service.WalletServiceInterface) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

// UpdateWallet
// @Accept json
// @Produce json
// @Summery Update Wallet
// @Tags Wallet
// @Param input body dto.UpdateWalletDTO true "operation info"
// @Router /wallet [post]
func (w *WalletHandler) UpdateWallet(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.UpdateWalletDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := w.walletService.UpdateWallet(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(200).JSON(response)
}

// GetWallet
// @Accept json
// @Produce json
// @Summery get wallet
// @Tags Wallet
// @Param WALLET_UUID path string true "UUID wallet" format(uuid)
// @Router /wallets/{WALLET_UUID} [get]
func (w *WalletHandler) GetWallet(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("WALLET_UUID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param WALLET_UUID exception", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, "invalid uuid")
	}

	wallet, err := w.walletService.GetWallet(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(wallet)
}
