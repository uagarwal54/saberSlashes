package main

import (
	"fmt"
	"hexagonalAppInGo/pkg/adding"
	"hexagonalAppInGo/pkg/http/rest"
	"hexagonalAppInGo/pkg/reading"
	"hexagonalAppInGo/pkg/storage"
	"log"
	"net/http"
)

func main() {
	storageObj, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("Error while connecting to DB: ", err)
	}
	// Here readingService is of the type service struct defined in reading package's service.go
	readingService := reading.NewService(storageObj)
	addingService := adding.NewService(storageObj)

	fmt.Println("Starting our server at port 8080...")
	router := rest.InitHandlers(readingService, addingService)
	log.Fatal(http.ListenAndServe(":8080", router))
}
