package main

import (
	"net/http"
)

func registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/store", handleIncomingMessage)
	router.HandleFunc("/getVal", handleGetValRequest)
}
