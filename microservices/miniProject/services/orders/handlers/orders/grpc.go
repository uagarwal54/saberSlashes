package handler

import (
	"cloudKitchen/services/common/genproto/orders"
	"cloudKitchen/services/orders/types"
	"context"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpcServer *grpc.Server, orderService types.OrderService) *OrdersGrpcHandler {
	grpcHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpcServer, grpcHandler)

	return grpcHandler
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:   42,
		CutomerID: 3,
		ProductID: 23,
		Quantity:  19,
	}
	if err := h.orderService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Success",
	}
	return res, nil
}
