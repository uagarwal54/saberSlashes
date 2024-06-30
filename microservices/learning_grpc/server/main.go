package main

import (
	pb "learningGRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	// RegisterGreetServiceServer is one of the auto generated functions in the  greet_grp.pb.go file which was auto generated from cli
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at: %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the grpc server: %v", err)
	}
}
