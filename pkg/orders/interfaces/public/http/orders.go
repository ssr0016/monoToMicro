package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/ssr0016/monolithToMicro/pkg/orders/application"
	"github.com/ssr0016/monolithToMicro/pkg/orders/domain/orders"
)

type orderResource struct {
	service    application.OrdersService
	repository orders.Repository
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderRequest `json:"address"`
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrderResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := orderResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

func (o orderResource) GetPaid(w http.ResponseWriter, r *http.Request) {

}
func (o orderResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadtRequest(err))
	}

	cmd := application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.NewV1().String()),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address),
	}

	if err := o.service.PlaceOrder(cmd); err != nil {
		render.JSON
	}

	w.WriteHeader(http.StatusOk)
	render.JSON(w, r, PostOrderResponse{
		OrderID: string(cmd.OrderID),
	})

}

func (o orderResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrNotFound(err))
		return
	}

	render.Respond(w, r, OrderPaidView{string(order.ID()), order.IsPaid()})
}
