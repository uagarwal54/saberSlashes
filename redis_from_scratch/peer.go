package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/tidwall/resp"
)

type (
	Peer struct {
		conn  net.Conn
		msgCh chan Message
	}

	Command    interface{}
	SetCommand struct {
		key, value []byte
	}
	GetCommand struct {
		key []byte
	}
)

const (
	CommandSet = "SET"
	CommandGet = "GET"
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
	rd := resp.NewReader(p.conn)
	for {
		values, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading value: ", err)
		}
		var cmd Command
		if values.Type() == resp.Array {
			for _, val := range values.Array() {
				switch val.String() {
				case CommandSet:
					if len(values.Array()) != 3 {
						return fmt.Errorf("invalid Number of variables for the SET command")
					}
					cmd = SetCommand{
						key:   values.Array()[1].Bytes(),
						value: values.Array()[2].Bytes(),
					}
					p.msgCh <- Message{cmd: cmd, peer: p}
					// return nil
				case CommandGet:
					if len(values.Array()) != 2 {
						return fmt.Errorf("invalid Number of variables for the GET command")
					}
					cmd = GetCommand{
						key: values.Array()[1].Bytes(),
					}
					p.msgCh <- Message{cmd: cmd, peer: p}
					// return nil
				default:

				}
			}
		}
	}
	return nil

}
