package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// keyValStore is initialized in common.go
	keyValStore.keyValStore = make(map[string]string)
	fmt.Println("Creating http router")
	router := http.NewServeMux()

	fmt.Println("Registering Routes")
	registerRoutes(router)
	go handleInterupt()
	go keyValStore.storeDataToDisk(triggerChannel)
	fmt.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8081", router))
}
