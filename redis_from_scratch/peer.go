package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/tidwall/resp"
)

type (
	// Peer struct has the peer -> server conn and the msg channel
	Peer struct {
		conn  net.Conn
		msgCh chan Message
		delCh chan *Peer
	}

	// Command is the interface which will represent all the reids commands
	Command interface{}

	// SetCommand is the most basic set command for key:val pair
	SetCommand struct {
		key, value []byte
	}
	// GetCommand is the most basic get command for key:val pair
	GetCommand struct {
		key []byte
	}

	HelloCommand struct {
		value string
	}
)

const (
	commandSet   = "SET"
	commandGet   = "GET"
	commandHello = "hello"
)

// NewPeer creates the Peer struct that is used top manage the peers in the server
func NewPeer(conn net.Conn, msgCh chan Message, delCh chan *Peer) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}

func (p *Peer) send(val []byte) error {
	_, err := p.conn.Write(val)
	return err
}

func (p *Peer) readLoop() error {
	// resp is the protocol that redis uses
	rd := resp.NewReader(p.conn)
	for {
		values, _, err := rd.ReadValue()
		if err == io.EOF {
			p.delCh <- p
			break
		}
		if err != nil {
			log.Fatal("Error while reading value: ", err)
		}
		var cmd Command
		if values.Type() == resp.Array {
			for _, val := range values.Array() {
				switch val.String() {
				case commandSet:
					if len(values.Array()) != 3 {
						return fmt.Errorf("invalid Number of variables for the SET command")
					}
					cmd = SetCommand{
						key:   values.Array()[1].Bytes(),
						value: values.Array()[2].Bytes(),
					}
					// return nil
				case commandGet:
					if len(values.Array()) != 2 {
						return fmt.Errorf("invalid Number of variables for the GET command")
					}
					cmd = GetCommand{
						key: values.Array()[1].Bytes(),
					}
					// return nil
				case commandHello:
					cmd = HelloCommand{
						value: values.Array()[1].String(),
					}

				default:
				}
				p.msgCh <- Message{cmd: cmd, peer: p}
			}
		}
	}
	return nil

}
