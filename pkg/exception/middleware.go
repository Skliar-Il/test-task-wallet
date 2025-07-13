package exception

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func Middleware(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	var fiberErr *fiber.Error
	var appErr *AppException
	switch {
	case errors.As(err, &appErr):
		code = appErr.Code
		message = appErr.Message
	case errors.As(err, &fiberErr):
		code = fiberErr.Code
		message = fiberErr.Message
	default:
		if err != nil {
			message = err.Error()
		}
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": message,
	})
}
