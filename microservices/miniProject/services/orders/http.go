package main

import (
	handler "cloudKitchen/services/orders/handlers/orders"
	"cloudKitchen/services/orders/service"
	"log"
	"net/http"
)

type (
	httpServer struct {
		addr string
	}
)

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("Waiting for orders on: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
