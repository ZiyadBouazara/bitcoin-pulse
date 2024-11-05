package domain

import (
	"time"
)

type PriceEvent struct {
	Type        string
	Sequence    int64
	ProductID   string
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
