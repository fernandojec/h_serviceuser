package users

import (
	"context"

	"github.com/fernandojec/h_serviceuser/config"

	"github.com/gofiber/fiber/v2"
)

type serviceUsers interface {
	saveRegistrasiService
}

type saveRegistrasiService interface {
	saveRegistrasi(ctx context.Context, req tableUsersModel, baseVerifyEmail string) (err error)
}

type handlerusers struct {
	svcUsers serviceUsers
}

func NewHandler_Users(svcusers serviceUsers) handlerusers {
	return handlerusers{svcUsers: svcusers}
}

func (h *handlerusers) saveRegistrasi(c *fiber.Ctx) error {

	req := new(tableUsersModel)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	// if err := ValidateStruct(req); err != nil {
	// 	// return writeErrorResponse(c, echo.ErrBadRequest.Code, err.Error())
	// 	return c.JSON(err)
	// }

	//save
	err := h.svcUsers.saveRegistrasi(
		c.UserContext(),
		*req,
		config.AppConfig.App.BaseUrl+config.AppConfig.App.BasePort+"/v1/auth/verify-email/",
	)

	if err != nil {
		// return writeErrorResponse(c, echo.ErrInternalServerError.Code, err.Error())
		if err.Error() == "email not sended" {
			return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
				"status":  fiber.StatusCreated,
				"message": "User has been created. But email fail to send",
			})
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Check your email to activate your account.",
	})
}
