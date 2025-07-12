package render

import "github.com/gofiber/fiber/v3"

func Error(err *fiber.Error, message string) error {
	if message != "" {
		err.Message = message
	}
	return err
}
