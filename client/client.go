package client

import (
	"github.com/oikomi/rtmp_server/handshake"
	"io"
)

type Client struct {
	Conn io.ReadWriter
}

func New(conn io.ReadWriter) (client *Client) {
	client = &Client{
		Conn: conn,
	}
	return
}

func (c *Client) Handshake() error {
	if err := handshake.With(&handshake.Param{
		Conn: c.Conn,
	}).Handshake(); err != nil {
		return err
	}
	return nil
}
