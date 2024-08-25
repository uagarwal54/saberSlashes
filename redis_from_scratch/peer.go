package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/tidwall/resp"
)

const (
	commandSet    = "SET"
	commandGet    = "GET"
	commandHello  = "HELLO"
	commandClient = "CLIENT"
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

	ClientCommand struct {
		value string
	}
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

		if values.Type() == resp.Array {
			rawCmd := values.Array()[0].String()
			rawCmd = strings.ToUpper(rawCmd)
			var cmd Command
			switch rawCmd {
			case commandClient:
				cmd = ClientCommand{
					value: values.Array()[1].String(),
				}

			case commandGet:
				if len(values.Array()) != 2 {
					return fmt.Errorf("invalid Number of variables for the GET command")
				}
				cmd = GetCommand{
					key: values.Array()[1].Bytes(),
				}

			case commandSet:
				if len(values.Array()) != 3 {
					return fmt.Errorf("invalid Number of variables for the SET command")
				}
				cmd = SetCommand{
					key:   values.Array()[1].Bytes(),
					value: values.Array()[2].Bytes(),
				}

			case commandHello:
				// This is the case where we are sending hello from some client which is NOT official redis client
				if len(values.Array()) == 1 {
					cmd = HelloCommand{
						value: values.Array()[0].String(),
					}

				} else {
					// This is handling the hello from the official client
					cmd = HelloCommand{
						value: values.Array()[1].String(),
					}
				}
			default:
				fmt.Printf("Got unknown command from client => %v\n", rawCmd)
			}
			p.msgCh <- Message{cmd: cmd, peer: p}
		}
	}
	return nil

}
