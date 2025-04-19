package connectors

import "net"

type (

	// Connector interface defines the methods that any connector should implement.
	Connector interface {
		Start() error
		HandleConnection(any)
		ReadLoop(any)

		// Getters for the connector
		GetListenAddr() string
		GetQuitChannel() *chan bool
		GetMessageChannel() *chan Message
		GetTCPListener() net.Listener
	}

	// Message struct will have all the info of the message, it will be agnostic of the connection method
	Message struct {
		From    string
		Payload []byte
	}
)
