package tcpconnector

import (
	"chatiyana/connectors"
	"fmt"
	"net"
)

type (
	// TCPServer struct will have all the info of the TCP server
	TCPServer struct {
		listenAddr        string
		TCPServerListener net.Listener
		Quitch            chan bool
		Msgch             chan connectors.Message
		Conns             map[string]*net.Conn
	}
)

// NewTCPServer initializes a new TCP server with the given address and a message channel
func NewTCPServer(listenAddr string) *TCPServer {
	return &TCPServer{
		listenAddr: listenAddr,
		Quitch:     make(chan bool),
		Msgch:      make(chan connectors.Message, 10),
		Conns:      make(map[string]*net.Conn),
	}
}

// Start starts the TCP server and listens for incoming connections
func (s *TCPServer) Start() error {
	var err error
	// Creates a listener on the given address and port and calls the handleConnection method to start the accept loop
	s.TCPServerListener, err = net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	fmt.Println("TCPServer started on", s.listenAddr)
	defer s.TCPServerListener.Close()
	s.HandleConnection(nil)
	<-s.Quitch
	close(s.Msgch)
	return nil
}

// HandleConnection handles the incoming connection and creates a new connection object
// Here there was no need to pass a parameter to the function as we are not using it but we are using it to implement the connector interface
func (s *TCPServer) HandleConnection(nothing any) {
	for {
		// Accepts any incoming connections and creates a connection object to handle the connection
		conn, err := s.TCPServerListener.Accept()
		if err != nil {
			return
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		s.Conns[conn.RemoteAddr().String()] = &conn
		go s.ReadLoop(conn)
	}
}

// ReadLoop reads the incoming messages from the connection and sends it to the message channel
func (s *TCPServer) ReadLoop(input any) {
	conn := input.(net.Conn)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}
		s.Msgch <- connectors.Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}

		conn.Write([]byte("Message recieved!!\n"))
	}
}

// GetListenAddr returns the listening address of the TCP server
func (s *TCPServer) GetListenAddr() string {
	return s.listenAddr
}

// GetQuitChannel returns the quit channel of the TCP server
func (s *TCPServer) GetQuitChannel() *chan bool {
	return &s.Quitch
}

// GetMessageChannel returns the message channel of the TCP server
func (s *TCPServer) GetMessageChannel() *chan connectors.Message {
	return &s.Msgch
}

// GetTCPListener returns the active connections of the TCP server,(This is not a good practice but we are doing it for practice)
func (s *TCPServer) GetTCPListener() net.Listener {
	return s.TCPServerListener
}

// func main() {
// 	server := NewServer(":3000")
// 	go func() {
// 		for msg := range server.msgch {
// 			payload := string(msg.payload)
// 			payload = strings.TrimSpace(payload)
// 			fmt.Printf("Received message from %s: %s\n", msg.from, payload)
// 		}
// 	}()
// 	log.Fatal(server.Start())
// }
