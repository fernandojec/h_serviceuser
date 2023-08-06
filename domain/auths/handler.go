package auths

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	iservice iservice
}

type iservice interface {
	SignIn(ctx context.Context, data signInRequest) (dataResponse signInRequestResponse, err error)
}

func NewAuthHandler(iservice iservice) *AuthHandler {
	return &AuthHandler{
		iservice: iservice,
	}
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	u := new(signInRequest)

	if err := c.BodyParser(u); err != nil {
		return err
	}
	errors := ValidateStructAuth(*u)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	dataReponse, err := h.iservice.SignIn(c.Context(), *u)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data":    dataReponse,
		"message": "success",
	})
}
