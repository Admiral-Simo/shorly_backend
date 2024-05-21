package handlers

import "github.com/gofiber/fiber/v2"

var (
	ErrInvalidRequestBody = fiber.Map{"err": "invalid request body"}
	ErrUnauthorized       = fiber.Map{"err": "incorrect username or password"}
)

func SendError(c *fiber.Ctx, statusCode int, err fiber.Map) error {
	return c.Status(statusCode).JSON(err)
}
