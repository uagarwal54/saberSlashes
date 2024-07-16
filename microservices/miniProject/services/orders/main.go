package main

import "log"

func main() {
	httpSever := NewHttpServer(":8000")
	go httpSever.Run()

	grpcServer := NewGRPCServer(":9000")
	log.Fatal(grpcServer.Run())
}
