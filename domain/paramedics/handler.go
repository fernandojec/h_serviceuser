package paramedics

import (
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	igrcp ParamedicClient
}

func NewHandler(igrcp ParamedicClient) *handler {
	return &handler{
		igrcp: igrcp,
	}
}

func (h *handler) CreateParamedics(c *fiber.Ctx) error {
	u := new(ParamedicCreateProto)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	dataCreated, err := h.igrcp.CreateParamedic(c.UserContext(), u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusCreated, &dataCreated)
}
func (h *handler) FindByHospital(c *fiber.Ctx) error {
	hospital := c.Params("hospital")

	if hospital == "" {
		return fiber.ErrBadRequest
	}
	u := ParamedicFindByHospitalProto{
		Hospitalid: hospital,
	}

	dataParamedic, err := h.igrcp.FindByHospital(c.UserContext(), &u)

	if err != nil {
		return err
	}

	return ifiber.SuccessResponse(c, fiber.StatusOK, &dataParamedic)
}
