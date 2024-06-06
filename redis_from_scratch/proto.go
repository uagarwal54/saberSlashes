package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSet = "SET"
)

type (
	Command    interface{}
	SetCommand struct {
		key, value string
	}
)

func parseCommand(raw string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))
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
						return nil, fmt.Errorf("invalid Number of variables for the set command")
					}
					cmd = SetCommand{
						key:   values.Array()[1].String(),
						value: values.Array()[2].String(),
					}
					return cmd, nil
				default:

				}
			}
		}
	}
	return nil, fmt.Errorf("invalid or unknown command recieved: %s", raw)
}
