package hospital

import (
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	svc HospitalServiceClient
}

func NewHandler(svc HospitalServiceClient) *handler {
	return &handler{svc: svc}
}

func (h *handler) List(c *fiber.Ctx) error {
	data_ret, err := h.svc.List(c.UserContext(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_ret)
}

func (h *handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	data_req := HospitalGetProto{
		HealthcareId: id,
	}
	data_ret, err := h.svc.Get(c.UserContext(), &data_req)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_ret)
}
func (h *handler) Create(c *fiber.Ctx) error {
	var data_req HospitalProto
	if err := c.BodyParser(&data_req); err != nil {
		return err
	}
	data_req.UserCreate = c.UserContext().Value(ifiber.USERID).(string)
	data_req.IsActive = true
	u := HospitalAddProto{
		Addhospital: &data_req,
	}
	data_ret, err := h.svc.Add(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_ret)
}
