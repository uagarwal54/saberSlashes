package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"redis_from_scratch/client"
	"time"
)

// start from 33.13
const defaultListenAddr = ":5001"

type (
	Config struct {
		ListenAddress string
	}
	Server struct {
		Config
		peers     map[*Peer]bool
		ln        net.Listener
		addPeerCh chan *Peer
		quitCh    chan struct{}
		msgCh     chan []byte
	}
)

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddress) == 0 {
		cfg.ListenAddress = defaultListenAddr
	}
	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan []byte),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		log.Println("Error while starting server")
		return err
	}
	s.ln = ln
	go s.loop()
	log.Println("Server Running, Listening at: ", s.ListenAddress)
	return s.acceptLoop()

}

// loop functin will add new peers if the redis is in cluster mode
func (s *Server) loop() {
	for {
		select {
		case rawMsg := <-s.msgCh: // rawMsg has the message bytes sent by the peer
			if err := s.handleRawMsg(rawMsg); err != nil {
				log.Print("Error while processing the raw mwssage: ", err)
			}
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		case <-s.quitCh:
			return
		}
	}
}

func (s *Server) handleRawMsg(rawMsg []byte) error {
	cmd, err := parseCommand(string(rawMsg))
	if err != nil {
		return err
	}
	// Doing v := cmd.(type) will create v with whatever the underlying type of the cmd is, since cmd is an interface so it will be of various command types of redis
	switch v := cmd.(type) {
	case SetCommand:
		fmt.Printf("Somebody wants to set the key: %s in the has table with the value: %s\n", v.key, v.value)
	}
	return nil
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println("Error while accepting the request: ", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn, s.msgCh)
	s.addPeerCh <- peer
	log.Println("New Peer added, remoteAddress: ", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		log.Println("Peer read error: err", err.Error()+" remoteAddr", conn.RemoteAddr())
	}
}

func main() {
	go func() {
		server := NewServer(Config{})
		log.Fatal(server.Start())
	}()
	time.Sleep(1 * time.Second)
	c := client.NewClient("localhost:5001")

	/* The diff between context.Background() and context.TODO() is nothing as both return context.emptyCtx which is an empty struct.
	The only diff that can be seen is that the context.Background() will return context.backgroundCtx struct which inherits context.emptyCtx and hence is an empty struct. context.backgroundCtx implements
	the string interface which returns "context.Background" as string
	whereas context.TODO() will return context.todoCtx struct which inherits context.emptyCtx and hence is an empty struct. context.todoCtx implements the string interface which returns "context.TODO" as string
	*/
	if err := c.Set(context.Background(), "foo", "bar"); err != nil {
		log.Fatal("Error: ", err)
	}
	select {} // We are blocking so that the program does not exit!
}
