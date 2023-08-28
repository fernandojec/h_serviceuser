package hospital

import (
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	svc          HospitalServiceClient
	app_newRelic *newrelic.Application
}

func NewHandler(svc HospitalServiceClient, _app_newRelic *newrelic.Application) *handler {
	return &handler{svc: svc, app_newRelic: _app_newRelic}
}

func (h *handler) List(c *fiber.Ctx) error {

	// fmt.Printf("%+v", *h.app_newRelic)
	// txn := h.app_newRelic.StartTransaction("main", nil, nil)
	// txn := newrelic.FromContext(c.Context())
	// txn := fibernewrelicmiddleware.Transaction(c.UserContext())
	// // if txn == nil {
	// // 	fmt.Println("no transaction")
	// // }
	// // txn.SetName("getListHealthcare")
	// defer txn.End()
	// ctx := newrelic.NewContext(context.Background(), txn)
	// if newrelic.FromContext(c.UserContext()) != nil {
	// 	fmt.Println("context already")
	// } else {
	// 	fmt.Println("no context")
	// }
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
