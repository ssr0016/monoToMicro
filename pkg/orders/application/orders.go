package application

type productsService interface {
}

type paymentService interface {
}

type OrdersService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommand struct {
}

// Struct Method
func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {

}
