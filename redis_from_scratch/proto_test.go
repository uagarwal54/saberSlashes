package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	// This is the format in which the redis resp protocol expects data, refer doc for resp
	raw := "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$3\r\nbar\r\n"
	cmd, err := parseCommand(raw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cmd)
}
