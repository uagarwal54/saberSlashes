package main

import (
	"chatiyana/connectors"
	tcpconnector "chatiyana/connectors/tcpConnector"
	websocketconnector "chatiyana/connectors/websocketConnector"
	"chatiyana/helpers"
	"fmt"
	"log"
	"strings"
)

func init() {
	if err := helpers.LoadConfigs(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	var server connectors.Connector

	if helpers.GetServerProtocol() == "tcp" {
		// Initialize the TCP server
		server = tcpconnector.NewTCPServer(helpers.GetServerAddress())
		go func(server connectors.Connector) {
			if err := server.Start(); err != nil {
				log.Fatalf("Error starting TCP server: %v", err)
			}
		}(server)
	} else if helpers.GetServerProtocol() == "websocket" {
		// Initialize the WebSocket server
		server = websocketconnector.NewWebsocketServer(helpers.GetServerAddress())
		go func(server connectors.Connector) {
			if err := server.Start(); err != nil {
				log.Fatalf("Error starting WebSocket server: %v", err)
			}
		}(server)
	}
	// The server.Start are called with in a go routine so that the main thread can continue to run and reach this point as boith the Start functions contain servers inside them
	messageChannel := server.GetMessageChannel()
	for msg := range *messageChannel {
		payload := strings.TrimSpace(string(msg.Payload))
		fmt.Printf("Received message from %s: %s\n", msg.From, payload)
	}

}
