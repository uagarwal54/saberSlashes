package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"
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

	// Command is the interface which will represent all the redis commands
	Command interface{}

	// SetCommand is the most basic set command for key:val pair
	SetCommand struct {
		key, value string
	}
	// GetCommand is the most basic get command for key:val pair
	GetCommand struct {
		key string
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
	reader := bufio.NewReader(p.conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		args := strings.Split(strings.TrimSpace(line), " ")
		argsCount := len(args)
		if argsCount > 0 {
			var (
				cmd Command
				err error
			)

			rawCmd, args, argsCount := args[0], args[1:], argsCount-1
			switch rawCmd {
			case commandClient:
				cmd, err = handleCommandClient(args, argsCount)
			case commandSet:
				cmd, err = handleCommandSet(args, argsCount)
			case commandGet:
				cmd, err = handleCommandGet(args, argsCount)
			case commandHello:
				cmd = handleCommandHello(args, argsCount)
			default:
				slog.Error("got unknown command from the client", slog.String("command", rawCmd))
			}

			if err != nil {
				return err
			}

			p.msgCh <- Message{cmd: cmd, peer: p}
		}
	}
}

func handleCommandClient(args []string, argsCount int) (ClientCommand, error) {
	expectedArgsCount := 1
	if argsCount != expectedArgsCount {
		return ClientCommand{}, fmt.Errorf("error: invalid command arguments provided for %s: expected %d arguments", commandClient, expectedArgsCount)
	}

	value := args[1]
	return ClientCommand{
		value: value,
	}, nil
}

func handleCommandSet(args []string, argsCount int) (SetCommand, error) {
	minArgsCount := 2
	if argsCount < minArgsCount {
		return SetCommand{}, fmt.Errorf("error: invalid command arguments provided for %s: expected %d arguments", commandSet, minArgsCount)
	}

	key, value := args[0], strings.Join(args[1:], " ")
	return SetCommand{
		key:   key,
		value: value,
	}, nil
}

func handleCommandGet(args []string, argsCount int) (GetCommand, error) {
	expectedArgsCount := 1
	if argsCount != expectedArgsCount {
		return GetCommand{}, fmt.Errorf("error: invalid command arguments provided for %s: expected %d arguments", commandGet, expectedArgsCount)
	}

	key := args[0]
	return GetCommand{
		key: key,
	}, nil
}

func handleCommandHello(args []string, argsCount int) HelloCommand {
	value := commandHello // This is the case where we are sending hello from some client which is NOT official redis client
	if argsCount >= 1 {
		value = strings.Join(args[1:], " ") // This is handling the hello from the official client
	}

	return HelloCommand{
		value: value,
	}
}
