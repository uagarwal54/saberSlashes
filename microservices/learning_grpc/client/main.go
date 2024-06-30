package main

import (
	"flag"
	pb "learningGRPC/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	streamType := flag.String("streamType", "unary", "Type of stream needed")
	flag.Parse()

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to server: ", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Udbhav", "Shrishti", "Rini"},
	}
	if *streamType == "unary" {
		callSayHello(client)
	} else if *streamType == "serverStream" {
		callSayHelloServerStream(client, names)
	} else if *streamType == "clientStream" {
		callSayHelloClientSideStreaming(client, names)
	} // else if *streamType == "bidirectionalStream" {
	// 	callSayHelloBiDirectionalStreamingServer(client, names)
	// }

}
