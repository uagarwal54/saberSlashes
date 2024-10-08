package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/tidwall/resp"
)

// start from 33.13
const defaultListenAddr = ":5001"

type (
	Config struct {
		ListenAddress string
	}
	Message struct {
		cmd  Command
		peer *Peer
	}
	Server struct {
		Config
		peers     map[*Peer]bool
		listener  net.Listener
		addPeerCh chan *Peer
		quitCh    chan struct{}
		delCh     chan *Peer
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
	// Doing v := cmd.(type) will create v with whatever the underlying type of the cmd is, since cmd is an interface so it will be of various command types of redis
	switch v := msg.cmd.(type) {
	case ClientCommand:
		if err := resp.NewWriter(msg.peer.conn).WriteString("OK"); err != nil {
			return err
		}
	case SetCommand:
		if err := s.kv.Set(v.key, v.value); err != nil {
			return err
		}
		if err := resp.NewWriter(msg.peer.conn).WriteString("OK"); err != nil {
			return err
		}

	case GetCommand:
		val, ok := s.kv.Get(v.key)
		if !ok {
			return fmt.Errorf("key %s, not found", v.key)
		}
		if err := resp.NewWriter(msg.peer.conn).WriteString(string(val)); err != nil {
			return err
		}

	case HelloCommand:
		spec := map[string]string{
			"server":  "redis",
			"role":    "master",
			"version": "6.0.0",
			"proto":   "3",
			"mode":    "standalone",
		}
		if err := msg.peer.send(respWriteMap(spec)); err != nil {
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
	peer := NewPeer(conn, s.msgCh, s.delCh)
	s.addPeerCh <- peer
	log.Println("New Peer added, remoteAddress: ", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		log.Println("Peer action error: err", err.Error()+" remoteAddr", conn.RemoteAddr())
	}
}

func main() {
	listenAddr := flag.String("addr", defaultListenAddr, "The address that the server will listen to")
	flag.Parse()
	server := NewServer(Config{
		ListenAddress: *listenAddr,
	})
	go func() {
		log.Fatal(server.Start())
	}()
	select {} // We are blocking so that the program does not exit!
}
