package exception

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func Middleware(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else if err != nil {
		message = err.Error()
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": message,
	})
}
