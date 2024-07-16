package types

import (
	"cloudKitchen/services/common/genproto/orders"
	"context"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
