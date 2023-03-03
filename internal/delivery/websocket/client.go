package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dmitryburov/go-coinbase-socket/config"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/dmitryburov/go-coinbase-socket/internal/usecase"
	"github.com/dmitryburov/go-coinbase-socket/pkg/exchange"
	"github.com/dmitryburov/go-coinbase-socket/pkg/exchange/coinbase"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger"
	"golang.org/x/sync/errgroup"
	"net"
	"strings"
	"sync"
	"time"
)

type client struct {
	logger   logger.Logger
	conn     exchange.Manager
	uc       *usecase.Services
	products []string
	channels []string
}

// NewSocketClient init websocket client from delivery layout
func NewSocketClient(conn exchange.Manager, uc *usecase.Services, logger logger.Logger, cfg config.ExchangeConfig) (*client, error) {
	if len(cfg.Symbols) == 0 {
		return nil, errors.New("not found symbols for subscribes")
	}

	return &client{
		logger,
		conn,
		uc,
		cfg.Symbols,
		cfg.Channels,
	}, nil
}

// Run websocket listener
func (c *client) Run(ctx context.Context) error {

	var g = errgroup.Group{}
	var hMap = make(map[string]chan entity.Ticker)

	for _, symbol := range c.products {
		hMap[symbol] = make(chan entity.Ticker)
		// readers
		g.Go(func() error {
			return c.uc.Exchange.Tick(ctx, hMap[symbol])
		})
	}

	// subscribe to products
	sData, _ := json.Marshal(map[string]interface{}{
		"type":        "subscribe",
		"product_ids": c.products,
		"channels":    c.channels,
	})
	_, err := c.conn.WriteData(sData)
	if err != nil {
		return err
	}

	message, err := c.readMessage()
	if err != nil {
		c.logger.Error(err)
		return err
	}
	result, err := coinbase.ParseResponse(message)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	switch result.Type {
	case coinbase.Error:
		c.logger.Fatal(fmt.Sprintf("%s:%s", result.Message, result.Reason))
	case coinbase.Subscriptions:
		c.logger.Info(fmt.Sprintf("started subscription on products [%s]", strings.Join(c.products, ",")))
	}

	// writers
	for _, symbol := range c.products {
		g.Go(func() error {
			return c.responseReader(symbol, hMap)
		})
	}

	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}

// readMessage from socket response
func (c *client) readMessage() ([]byte, error) {
	var message = make([]byte, 1024) //TODO need change global?

	n, err := c.conn.ReadData(message)
	if err != nil {
		return nil, err
	}

	return message[:n], nil
}

// responseReader write to symbol channel from response socket data
func (c *client) responseReader(symbol string, hMap map[string]chan entity.Ticker) error {
	var mu = sync.Mutex{}
	var tickData *coinbase.Response

	for {
		message, err := c.readMessage()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			c.logger.Error(err)
			continue
		}

		tickData, err = coinbase.ParseResponse(message)
		if err != nil {
			c.logger.Error(err)
			continue
		}

		switch tickData.Type {
		case coinbase.Error:
			//TODO need break exchange if error?
			c.logger.Error(err)
			continue
		case coinbase.Ticker:
			mu.Lock()
			hMap[symbol] <- entity.Ticker{
				Timestamp: time.Now().UnixNano(), // for exclude collision and accuracy time of ticker
				Bid:       tickData.BestBid,
				Ask:       tickData.BestAsk,
				Symbol:    tickData.ProductID,
			}
			mu.Unlock()
		}
	}

	return nil
}
