package domain

import (
	"time"
)

type PriceEvent struct {
	Type        string
	Sequence    int64
	ProductID   Stock
	Price       float64
	Open24H     float64
	Volume24H   float64
	Low24H      float64
	High24H     float64
	Volume30D   float64
	BestBid     float64
	BestBidSize float64
	BestAsk     float64
	BestAskSize float64
	Side        string
	Time        time.Time
	TradeId     int64
	LastSize    float64
}

type SubscriptionMessage struct {
	Action Action `json:"action"`
	Stock  Stock  `json:"stock"`
}

type ErrorMessage struct {
	Type    string
	Message string
}

type Action string

const (
	Subscribe   Action = "subscribe"
	Unsubscribe Action = "unsubscribe"
)

type Stock string

const (
	StockBitcoin Stock = "BTC-USD"
)

func IsSupportedStock(stock string) bool {
	switch Stock(stock) {
	case StockBitcoin:
		return true
	default:
		return false
	}
}
