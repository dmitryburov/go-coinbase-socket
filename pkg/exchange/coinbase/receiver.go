package coinbase

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Type      ResponseType `json:"type,string"`
	Message   string       `json:"message,omitempty"`
	Reason    string       `json:"reason,omitempty"`
	ProductID string       `json:"product_id"`
	BestBid   float64      `json:"best_bid,string"`
	BestAsk   float64      `json:"best_ask,string"`
}

type ResponseType int

const (
	Error ResponseType = iota
	Subscriptions
	Heartbeat
	Ticker
	Level2
)

var responseTypeNames = [...]string{"error", "subscriptions", "heartbeat", "ticker", "level2"}

func (r ResponseType) String() string {
	return responseTypeNames[r]
}

func (r *ResponseType) UnmarshalJSON(v []byte) error {
	str := string(v)
	for i, name := range responseTypeNames {
		if name == str {
			*r = ResponseType(i)
			return nil
		}
	}

	return fmt.Errorf("invalid locality type %q", str)
}

func ParseResponse(message []byte) (response *Response, err error) {
	err = json.Unmarshal(message, &response)
	if err != nil {
		return nil, err
	}

	return
}
