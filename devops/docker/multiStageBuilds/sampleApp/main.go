package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
)

func main(){
	fmt.Println("This is to test the docker image only....")
	router := http.NewServeMux()
	router.HandleFunc("/store", handleIncomingMessage)
	fmt.Println("Starting the server, listening at: 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func handleIncomingMessage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.Header().Set("Status", "403")
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error while reading the request body", err)
		w.Header().Set("Status", "400")
		w.Header().Set("Message", "Error while reading the request body")
	}
	fmt.Println(string(bodyBytes))
}
