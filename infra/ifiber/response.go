package ifiber

import "github.com/gofiber/fiber/v2"

type Error struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}
type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(
		Error{
			Code:  statusCode,
			Error: data,
		},
	)
}
func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(
		Success{
			Code: statusCode,
			Data: data,
		},
	)
}
