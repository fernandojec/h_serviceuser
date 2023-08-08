package auths

import (
	"context"

	customvalidator "github.com/fernandojec/h_serviceuser/pkg/customValidator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	iservice iservice
}

type iservice interface {
	SignIn(ctx context.Context, data signInRequest) (dataResponse signInRequestResponse, err error)
	SignOut(ctx context.Context, data signOutRequest) (err error)
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
	validate := customvalidator.NewCustomValidator()
	errors := validate.Validate(u)
	if errors != nil {
		return errors
	}

	dataReponse, err := h.iservice.SignIn(c.UserContext(), *u)
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

func (h *AuthHandler) SignOut(c *fiber.Ctx) error {
	tokenString := c.Get("x-token")
	data := signOutRequest{
		AccessToken: tokenString,
	}
	err := h.iservice.SignOut(c.UserContext(), data)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("User logged out")
}
