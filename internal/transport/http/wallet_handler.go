package http

import (
	"github.com/Skliar-Il/test-task-wallet/internal/dto"
	"github.com/Skliar-Il/test-task-wallet/internal/service"
	"github.com/Skliar-Il/test-task-wallet/pkg/logger"
	"github.com/Skliar-Il/test-task-wallet/pkg/render"
	"github.com/gofiber/fiber/v3"
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

	response, err := w.walletService.UpdateWallet(ctx.Context(), body)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response)
}
