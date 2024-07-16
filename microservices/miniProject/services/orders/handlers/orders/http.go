package handler

import (
	"cloudKitchen/services/common/genproto/orders"
	"cloudKitchen/services/common/utils"
	"cloudKitchen/services/orders/types"
	"net/http"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		orderService: orderService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:   42,
		CutomerID: 3,
		ProductID: 23,
		Quantity:  19,
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := &orders.CreateOrderResponse{Status: "SUCCESS"}
	utils.WriteJSON(w, http.StatusOK, res)
}
