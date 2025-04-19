package helpers

// Package helpers provides utility functions and constants for the application.
import (
	"os"

	"github.com/joho/godotenv"
)

// serverHost is the address of the server
var serverHost string

// serverPort is the port on which the server will listen for incoming connections
var serverPort string

// serverProtocol is the protocol used by the server (tcp or websocket)
var serverProtocol string

// LoadConfigs loads the server configurations from environment variables or defaults to predefined values
func LoadConfigs() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000" // Default to port 8080 if not set
	}
	serverPort = port
	serverHost = os.Getenv("SERVER_ADDRESS")
	if serverHost == "" {
		serverHost = "localhost" // Default to localhost if not set
	}
	serverProtocol = os.Getenv("SERVER_PROTOCOL")
	if serverProtocol == "" {
		serverProtocol = "websocket" // Default to http if not set
	} else if serverProtocol != "tcp" && serverProtocol != "websocket" {
		serverProtocol = "websocket" // Default to websocket if not set
	}
	return nil
}

// GetServerHost returns the server host
func GetServerHost() string {
	return serverHost
}

// GetServerPort returns the server port
func GetServerPort() string {
	return serverPort
}

// GetServerProtocol returns the server protocol
func GetServerProtocol() string {
	return serverProtocol
}

// GetServerAddress returns the server address
func GetServerAddress() string {
	return serverHost + ":" + serverPort
}

// GetServerURL returns the server URL in the format protocol://address:port
func GetServerURL() string {
	return serverProtocol + "://" + serverHost + ":" + serverPort
}
