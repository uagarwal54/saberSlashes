package main

import (
	handler "cloudKitchen/services/orders/handlers/orders"
	"cloudKitchen/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{addr: addr}
}

func (s *grpcServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)
	log.Println("Starting the Order's internal GRPC server on: ", s.addr)
	return grpcServer.Serve(listener)
}
