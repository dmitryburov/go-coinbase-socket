package coinbase

import (
	"encoding/json"
	"time"
)

type Response struct {
	Type      string    `json:"type"`
	ProductID string    `json:"product_id"`
	BestBid   float64   `json:"best_bid,string"`
	BestAsk   float64   `json:"best_ask,string"`
	Time      time.Time `json:"time"`
}

type ResponseType int

const (
	Error ResponseType = iota
	Subscriptions
	Heartbeat
	Ticker
	Level2
)

func (r ResponseType) String() string {
	return [...]string{"error", "subscriptions", "heartbeat", "ticker", "level2"}[r]
}

func ParseResponse(message []byte) (response *Response, err error) {
	err = json.Unmarshal(message, &response)
	if err != nil {
		return nil, err
	}

	return
}
