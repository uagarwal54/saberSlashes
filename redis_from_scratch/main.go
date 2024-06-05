package main

import (
	"fmt"
	"log"
	"net"
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
				log.Fatal("Error while processing the raw mwssage: ", err)
			}
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		case <-s.quitCh:
			return
		}
	}
}

func (s *Server) handleRawMsg(rawMsg []byte) error {
	fmt.Println(string(rawMsg))
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
	server := NewServer(Config{})
	log.Fatal(server.Start())
}
