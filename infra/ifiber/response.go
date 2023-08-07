package ifiber

import "github.com/gofiber/fiber/v2"

type Error struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(
		Error{
			Code:  statusCode,
			Error: data,
		},
	)
}
