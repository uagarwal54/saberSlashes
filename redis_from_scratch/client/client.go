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

func (c *Client) Set(ctx context.Context, key string, val any) error {
	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	var respVal []resp.Value
	respVal = append(respVal, resp.StringValue("SET"))
	respVal = append(respVal, resp.StringValue(key))
	respVal = append(respVal, resp.AnyValue(val))
	wr.WriteArray(respVal)
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

func (c *Client) Close() error {
	return c.conn.Close()
}
