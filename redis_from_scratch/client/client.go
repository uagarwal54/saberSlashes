package client

import (
	"context"
	"fmt"
	"net"
)

type Client struct {
	addr string
	conn net.Conn
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Client{
		addr: address,
		conn: conn,
	}, nil
}

func (c *Client) Set(ctx context.Context, key, val string) error {
	data := fmt.Sprintf("SET %s %s\n", key, val)
	_, err := c.conn.Write([]byte(data))
	return err
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	data := fmt.Sprintf("GET %s\n", key)
	_, err := c.conn.Write([]byte(data))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	return string(buffer[:n]), err
}

func (c *Client) Close() error {
	return c.conn.Close()
}
