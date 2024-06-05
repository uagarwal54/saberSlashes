package main

import (
	"io"
	pb "learningGRPC/proto"
	"log"
)

// SayHelloClientSideStreaming handels the stream being sent from the client to the server, the servers job is to process the stream and send a response at the end of the stream
func (s *helloServer) SayHelloClientSideStreaming(stream pb.GreetService_SayHelloClientSideStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: ", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
	return nil
}
