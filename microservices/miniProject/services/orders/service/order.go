package service

// service will contain the business logic
import (
	"cloudKitchen/services/common/genproto/orders"
	"context"
)

var ordersDB = make([]*orders.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}
