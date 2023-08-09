package patient

import (
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	ipatientgrpc PatientProtoServiceClient
}

func NewHandler(ipatientgrpc PatientProtoServiceClient) *handler {
	return &handler{ipatientgrpc: ipatientgrpc}
}

func (h *handler) Add(c *fiber.Ctx) error {
	u := new(PatientProto)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	u.UserCreate = c.UserContext().Value(ifiber.USERID).(string)
	_, err := h.ipatientgrpc.Create(c.UserContext(), u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusCreated, nil)
}

func (h *handler) Find(c *fiber.Ctx) error {
	u := new(PatientFilterProto)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	data_result, err := h.ipatientgrpc.Find(c.UserContext(), u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_result)
}

func (h *handler) List(c *fiber.Ctx) error {
	data_result, err := h.ipatientgrpc.List(c.UserContext(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_result)
}
