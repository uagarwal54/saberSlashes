package rest

import (
	"hexagonalAppInGo/pkg/adding"
	"hexagonalAppInGo/pkg/reading"

	"github.com/gorilla/mux"
)

// Here the readingService object is defined as an instance of the Service interface defined in reading/service.go
func InitHandlers(readingService reading.Service, addingService adding.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", welcomeHanlder()).Methods("GET")
	router.HandleFunc("/api/getAllCandies/", getAllCandiesHandler(readingService)).Methods("GET")
	router.HandleFunc("/api/addNewCandies/", addCandyHandler(addingService)).Methods("POST")
	return router
}
