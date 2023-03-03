package entity

// Ticker model data of response from exchange
type Ticker struct {
	Timestamp int64
	Bid       float64
	Ask       float64
	Symbol    string
}
