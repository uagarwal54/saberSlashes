package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, dest interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading the body")
		return err
	}
	return json.Unmarshal(body, dest)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	http.Error(w, err.Error(), statusCode)
}

func WriteJSON(w http.ResponseWriter, statusCode int, res interface{}) {
	responseBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Error while Marshalling a response: ", err)
	}
	w.WriteHeader(statusCode)
	w.Write(responseBytes)
}
