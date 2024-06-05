package main

import (
	"fmt"
	pb "learningGRPC/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerSideStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerSideStreamingServer) error {
	log.Printf("got the req with names: %v", req.Names)
	for _, name := range req.Names {
		fmt.Println("Processing name: " + name)
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
