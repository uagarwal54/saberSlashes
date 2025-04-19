package websocketconnector

// Package websockets implements a simple WebSocket WebsocketServer in Go.

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"chatiyana/connectors"
)

// WebsocketServer struct will have all the info of the websocket WebsocketServer
type WebsocketServer struct {
	listenAddr string
	Conns      map[string]*websocket.Conn
	Msgch      chan connectors.Message
	Quitch     chan bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for testing)
	},
}

// NewWebsocketServer to initialize the WebsocketServer struct
func NewWebsocketServer(listenAddress string) *WebsocketServer {
	return &WebsocketServer{
		listenAddr: listenAddress,
		Conns:      make(map[string]*websocket.Conn),
		Msgch:      make(chan connectors.Message, 10),
		Quitch:     make(chan bool),
	}
}

// Start starts the TCP server and listens for incoming connections
func (s *WebsocketServer) Start() error {
	// WebsocketServer starts listening on the address
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}
		s.HandleConnection(conn)
	})
	fmt.Println("Websocket Server started on", s.listenAddr)
	http.ListenAndServe(s.listenAddr, nil)
	return nil
}

// HandleConnection handles the incoming connection and creates a new connection object
func (s *WebsocketServer) HandleConnection(connection any) {
	conn, _ := connection.(*websocket.Conn)
	// defer conn.Close()
	fmt.Println("New incoming connection from client: ", conn.RemoteAddr().String())
	var mutex sync.Mutex

	// We are using mutex lock here as maps are not concurrent safe in goalng
	mutex.Lock()
	s.Conns[conn.RemoteAddr().String()] = conn
	mutex.Unlock()
	go s.ReadLoop(conn)
}

// ReadLoop reads the incoming messages from the connection and sends it to the message channel
func (s *WebsocketServer) ReadLoop(input any) {
	fmt.Println("Reading...")
	ws := input.(*websocket.Conn)
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// This will close the connection from the WebsocketServer side if the connection on the client side has closed
			if err == io.EOF {
				fmt.Println("Connection closed by client: ", ws.RemoteAddr().String())
				ws.Close()
			} else {
				// Here we are continuing to read from the connection even if something wrong has been passed from the client
				// We can very well kill the connection too after logging the error, it is up to us to decide
				fmt.Println("Read Error: ", err)
				ws.Close()
				break
			}
		}
		// fmt.Println(string(msg))
		s.Msgch <- connectors.Message{
			From:    ws.RemoteAddr().String(),
			Payload: msg,
		}
		// s.broadCast(msg)
	}
}

func (s *WebsocketServer) broadCast(b []byte) {
	for addr := range s.Conns {
		go func(ws *websocket.Conn) {
			if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
				fmt.Println("Write Error: ", err)
			}
		}(s.Conns[addr])
	}
}

// GetListenAddr returns the listening address of the TCP server
func (s *WebsocketServer) GetListenAddr() string {
	return s.listenAddr
}

// GetQuitChannel returns the quit channel of the TCP server
func (s *WebsocketServer) GetQuitChannel() *chan bool {
	return &s.Quitch
}

// GetMessageChannel returns the message channel of the TCP server
func (s *WebsocketServer) GetMessageChannel() *chan connectors.Message {
	return &s.Msgch
}

// GetTCPListener returns the active connections of the TCP server,(This is not a good practice but we are doing it for practice)
func (s *WebsocketServer) GetTCPListener() net.Listener {
	return nil
}
