package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type (
	keyVal struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

func handleIncomingMessage(w http.ResponseWriter, r *http.Request) {
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
	var kv keyVal
	if err = json.Unmarshal(bodyBytes, &kv); err != nil {
		log.Fatal("Error while unmarshalling the request body", err)
		w.Header().Set("Status", "400")
		w.Header().Set("Message", "Error while unmarshalling the request body")
	}
	currentTime := time.Now().Format("2006-01-02T15:04:05")
	keyValStore.mu.Lock()
	keyValStore.keyValStore[kv.Key] = kv.Value + "!@!" + currentTime
	keyValStore.mu.Unlock()
	fmt.Println(keyValStore.keyValStore)
	w.Header().Set("Status", "200")
	w.Header().Set("Message", "Data stored successfully")
}

func handleGetValRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.Header().Set("Staus", "403")
	}
	key := r.URL.Query().Get("key")
	fmt.Println("Key is", key)
	keyValStore.mu.RLock()

	if _, ok := keyValStore.keyValStore[key]; ok {
		w.Header().Set("Status", "200")
		w.Header().Set("Message", "Data fetched successfully")
		w.WriteHeader(http.StatusOK)
		val := keyValStore.keyValStore[key]

		data := map[string]string{"key": key, "value": strings.Split(val, "!@!")[0]}
		fmt.Println(data)
		json.NewEncoder(w).Encode(data)
	} else {
		w.Header().Set("Status", "404")
		w.Header().Set("Message", "Data not found")
	}
	keyValStore.mu.RUnlock()
}
