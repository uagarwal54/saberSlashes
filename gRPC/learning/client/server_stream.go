package main

import (
	"context"
	"io"
	pb "learningGRPC/proto"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming has started")
	stream, err := client.SayHelloServerSideStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while streaming: ", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
