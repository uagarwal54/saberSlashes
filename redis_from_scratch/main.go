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
	Message struct {
		data []byte
		peer *Peer
	}
	Server struct {
		Config
		peers     map[*Peer]bool
		listener  net.Listener
		addPeerCh chan *Peer
		quitCh    chan struct{}
		msgCh     chan Message
		kv        KV
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
		msgCh:     make(chan Message),
		kv:        NewKV(),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		log.Println("Error while starting server")
		return err
	}
	s.listener = listener
	go s.startServerLoop()
	log.Println("Server Running, Listening at: ", s.ListenAddress)
	return s.acceptPeerLoop()

}

// loop functin will add new peers if the redis is in cluster mode as well as it will start the processing of the command coming in
func (s *Server) startServerLoop() {
	for {
		select {
		case msg := <-s.msgCh: // rawMsg has the message bytes sent by the peer
			if err := s.handleMsg(msg); err != nil {
				log.Print("Error while processing the raw mwssage: ", err)
			}
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		case <-s.quitCh:
			return
		}
	}
}

func (s *Server) handleMsg(msg Message) error {
	cmd, err := parseCommand(string(msg.data))
	if err != nil {
		return err
	}
	// Doing v := cmd.(type) will create v with whatever the underlying type of the cmd is, since cmd is an interface so it will be of various command types of redis
	switch v := cmd.(type) {
	case SetCommand:
		return s.kv.Set(v.key, v.value)
	case GetCommand:
		val, ok := s.kv.Get(v.key)
		if !ok {
			return fmt.Errorf("key not found")
		}
		if err = msg.peer.send(val); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) acceptPeerLoop() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("Error while accepting the request: ", err)
			continue
		}
		// Each peer connecting to the server will have an instance of handleConn GoR running for them
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn, s.msgCh)
	s.addPeerCh <- peer
	log.Println("New Peer added, remoteAddress: ", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		log.Println("Peer action error: err", err.Error()+" remoteAddr", conn.RemoteAddr())
	}
}

func main() {
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(1 * time.Second)

	c := client.NewClient("localhost:5001")
	for i := 1; i < 10; i++ {
		/* The diff between context.Background() and context.TODO() is nothing as both return context.emptyCtx which is an empty struct.
		The only diff that can be seen is that the context.Background() will return context.backgroundCtx struct which inherits context.emptyCtx and hence is an empty struct. context.backgroundCtx implements
		the string interface which returns "context.Background" as string
		whereas context.TODO() will return context.todoCtx struct which inherits context.emptyCtx and hence is an empty struct. context.todoCtx implements the string interface which returns "context.TODO" as string
		*/
		if err := c.Set(context.Background(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal("Error: ", err)
		}
		if val, err := c.Get(context.Background(), fmt.Sprintf("foo_%d", i)); err != nil {
			log.Fatal("Error: ", err)
		} else {
			fmt.Println("Value: ", val)
		}
	}
	// fmt.Println(server.kv.data)
	time.Sleep(1 * time.Second)

	select {} // We are blocking so that the program does not exit!
}
