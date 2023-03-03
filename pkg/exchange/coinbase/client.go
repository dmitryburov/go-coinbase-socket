package coinbase

import (
	"fmt"
	"golang.org/x/net/websocket"
)

const (
	ErrRequireConfigParameters = "not correct input parameters"
)

type Subscribe struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

type client struct {
	*websocket.Conn
}

// NewCoinbaseClient init client for Coinbase
func NewCoinbaseClient(url, protocol, origin string) (*client, error) {
	if origin == "" || url == "" {
		return nil, fmt.Errorf("%s", ErrRequireConfigParameters)
	}

	conn, err := websocket.Dial(url, protocol, origin)
	if err != nil {
		return nil, err
	}

	return &client{conn}, nil
}

func (c *client) WriteData(message []byte) (int, error) {
	return c.Write(message)
}

func (c *client) ReadData() ([]byte, error) {
	var message = make([]byte, 512) //TODO need change global? 1MB

	n, err := c.Read(message)
	if err != nil {
		return nil, err
	}

	return message[:n], nil
}

func (c *client) CloseConnection() error {
	return c.Close()
}
