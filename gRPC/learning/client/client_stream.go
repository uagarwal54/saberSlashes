package main

import (
	"context"
	pb "learningGRPC/proto"
	"log"
	"time"
)

func callSayHelloClientSideStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloClientSideStreaming(context.Background())
	if err != nil {
		log.Fatal("Error while sending the stream: ", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %s", err)
		}
		log.Fatalf("Send req with the name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while recieving the response from the server: ", err)
	}
	log.Println("Client Streaming finished")
	log.Printf("%v", res.Messages)
}
