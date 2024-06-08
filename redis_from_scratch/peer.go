package main

import (
	"net"
)

type (
	Peer struct {
		conn  net.Conn
		msgCh chan Message
	}
)

func NewPeer(conn net.Conn, msgCh chan Message) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}

func (p *Peer) send(val []byte) error {
	_, err := p.conn.Write(val)
	return err
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])

		p.msgCh <- Message{data: msgBuf, peer: p}
	}

}
