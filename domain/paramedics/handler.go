package paramedics

import "github.com/gofiber/fiber/v2"

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
	_, err := h.igrcp.CreateParamedic(c.UserContext(), u)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}
func (h *handler) FindByHospital(c *fiber.Ctx) error {
	hospital := c.Params("hospital")

	if hospital == "" {
		return fiber.ErrBadRequest
	}
	u := ParamedicFindByHospitalProto{
		Hospitalid: hospital,
	}

	_, err := h.igrcp.FindByHospital(c.UserContext(), &u)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
