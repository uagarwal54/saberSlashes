package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// keyValStore is initialized in common.go
	keyValStore.keyValStore = make(map[string]string)
	fmt.Println("Creating http router")
	if _, err := os.Stat("./store"); os.IsNotExist(err) {
		if err = os.Mkdir("./store", 0750); err != nil {
			log.Fatal("Unable to create store dir: ", err)
		}
	}
	router := http.NewServeMux()

	fmt.Println("Registering Routes")
	registerRoutes(router)
	go handleInterupt()
	go keyValStore.storeDataToDisk(triggerChannel)
	fmt.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8081", router))
}
