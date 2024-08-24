package client

import (
	"bytes"
	"context"
	"net"

	"github.com/tidwall/resp"
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
	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	wr.WriteArray(
		[]resp.Value{resp.StringValue("SET"),
			resp.StringValue(key),
			resp.StringValue(val)},
	)
	_, err := c.conn.Write(buf.Bytes())
	return err
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	wr.WriteArray(
		[]resp.Value{resp.StringValue("GET"),
			resp.StringValue(key),
		})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	returnBuf := make([]byte, 1024)
	n, err := c.conn.Read(returnBuf)
	return string(returnBuf[:n]), err
}
