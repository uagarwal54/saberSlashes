package main

import (
	"context"
	pb "learningGRPC/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Could not greet: ", err)
	}

	log.Printf("%s", res.Message)
}
