package rest

import (
	"encoding/json"
	"fmt"
	"hexagonalAppInGo/pkg/adding"
	"hexagonalAppInGo/pkg/reading"
	"net/http"
	"strconv"
)

func welcomeHanlder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our Candy Shop")
	}
}

func getAllCandiesHandler(readingService reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		candieNames, err := readingService.GetAllCandies()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Cannot process the request", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(candieNames)
	}
}

func addCandyHandler(addingService adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCandy adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&newCandy); err != nil {
			fmt.Println(err)
			http.Error(w, "Error while UnMarshling Request body of the request to add a new candy", http.StatusInternalServerError)
			return
		}
		id, err := addingService.AddCandy(newCandy)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Cannot process the request", http.StatusInternalServerError)
			return
		}
		newCandy.Id, _ = strconv.Atoi(id)
		json.NewEncoder(w).Encode("Added the new candy")
	}
}
