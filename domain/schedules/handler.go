package schedules

import (
	"strconv"

	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	ischedulegrpc ScheduleClient
}

func NewHandler(ischedulegrpc ScheduleClient) *handler {
	return &handler{ischedulegrpc: ischedulegrpc}
}

func (h *handler) CreateSchedule(c *fiber.Ctx) error {
	dataBody := new(RequestAdd)
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}
	u := dataBody.ConvertToScheduleCreateProto()
	u.Usercreate = c.UserContext().Value(ifiber.USERID).(string)
	dataResult, err := h.ischedulegrpc.CreateSchedule(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusCreated, dataResult)
}

func (h *handler) FindScheduleAll(c *fiber.Ctx) error {
	dataBody := new(RequestFindSchedule)
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}
	u := dataBody.ConvertToScheduleFindProto()
	dataResult, err := h.ischedulegrpc.FindScheduleAll(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, dataResult)
}

func (h *handler) FindScheduleAvailable(c *fiber.Ctx) error {
	dataBody := new(RequestFindSchedule)
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}
	u := dataBody.ConvertToScheduleFindProto()
	dataResult, err := h.ischedulegrpc.FindScheduleAvailable(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, dataResult)
}

func (h *handler) FindScheduleBooked(c *fiber.Ctx) error {
	dataBody := new(RequestFindSchedule)
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}
	u := dataBody.ConvertToScheduleFindProto()
	dataResult, err := h.ischedulegrpc.FindScheduleBooked(c.UserContext(), &u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, dataResult)
}

func (h *handler) BookSlot(c *fiber.Ctx) error {
	u := new(SetScheduleProto)
	scheduleId := c.Params("schedule_id")
	if scheduleId == "" {
		return ifiber.ErrorResponse(c, fiber.StatusBadRequest, "schedule_id is required")
	}
	scheduleIdInt, err := strconv.ParseInt(scheduleId, 10, 32)
	if err != nil {
		return ifiber.ErrorResponse(c, fiber.StatusBadRequest, "schedule_id is invalid")
	}
	u.Scheduleslotid = int32(scheduleIdInt)
	dataResult, err := h.ischedulegrpc.BookSlot(c.UserContext(), u)
	if err != nil {
		return err
	}
	return ifiber.SuccessResponse(c, fiber.StatusOK, dataResult)
}
