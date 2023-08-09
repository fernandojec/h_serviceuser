package appointment

import (
	"errors"

	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	iappointmentsvc AppointmentServiceClient
}

func NewHandler(iappointmentsvc AppointmentServiceClient) *handler {
	return &handler{iappointmentsvc: iappointmentsvc}
}

// func (h *handler) List(c *fiber.Ctx) error {
// 	u := new(AppointmentProto)

// 	u.AppointmentTime=c.Query("")
// 	// if err := c.QueryParser(u); err!= nil {
//     //     return c.Status(400).JSON(err)
//     // }
// 	h.iappointmentsvc.List(c.UserContext(),*u)
// }

func (h *handler) Get(c *fiber.Ctx) error {
	appointment_no := c.Params("appointment_no")
	healthcare_id := c.Params("healthcare_id")
	if appointment_no == "" || healthcare_id == "" {
		return errors.New("missing parameter appointment_no or healthcare_id")
	}
	data := AppointmentGetProto{
		HealthcareId:  healthcare_id,
		AppointmentNo: appointment_no,
	}
	data_result, err := h.iappointmentsvc.Get(c.UserContext(), &data)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_result)
}

func (h *handler) Add(c *fiber.Ctx) error {
	// u := new(AppointmentProto)
	var data RequestAdd
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	u := data.ConvertToAppointmentProto()
	u.IsVoid = false
	u.UserCreate = c.UserContext().Value(ifiber.USERID).(string)
	data_req := AppointmentAddProto{
		Addappointment: &u,
	}
	data_result, err := h.iappointmentsvc.Add(c.UserContext(), &data_req)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusCreated, data_result)
}

func (h *handler) Update(c *fiber.Ctx) error {
	appointment_no := c.Params("appointment_no")
	healthcare_id := c.Params("healthcare_id")
	if appointment_no == "" || healthcare_id == "" {
		return errors.New("missing parameter appointment_no or healthcare_id")
	}
	data_body := new(RequestEdit)
	if err := c.BodyParser(data_body); err != nil {
		return err
	}
	u := data_body.ConvertToAppointmentProto()
	u.AppointmentNo = appointment_no
	u.HealthcareId = healthcare_id

	u.UserCreate = c.UserContext().Value(ifiber.USERID).(string)
	data_req := AppointmentUpdateProto{
		Updateappointment: &u,
	}
	data_result, err := h.iappointmentsvc.Update(c.UserContext(), &data_req)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_result)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	appointment_no := c.Params("appointment_no")
	healthcare_id := c.Params("healthcare_id")
	if appointment_no == "" || healthcare_id == "" {
		return errors.New("missing parameter appointment_no or healthcare_id")
	}
	u := AppointmentDeleteProto{
		Deleteappointment: &AppointmentProto{
			AppointmentNo: appointment_no,
			HealthcareId:  healthcare_id,
		},
	}
	data_result, err := h.iappointmentsvc.Delete(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, data_result)
}
