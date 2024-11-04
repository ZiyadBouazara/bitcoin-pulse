package domain

import (
	"fmt"
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

func (event *PriceEvent) String() string {
	return fmt.Sprintf("Type: %s, Sequence: %d, ProductID: %s, Price: %f, Open24H: %f, Volume24H: %f, Low24H: %f, High24H: %f, Volume30D: %f, BestBid: %f, BestBidSize: %f, BestAsk: %f, BestAskSize: %f, Side: %s, Time: %s, TradeId: %d, LastSize: %f",
		event.Type, event.Sequence, event.ProductID, event.Price, event.Open24H, event.Volume24H, event.Low24H, event.High24H, event.Volume30D, event.BestBid, event.BestBidSize, event.BestAsk, event.BestAskSize, event.Side, event.Time, event.TradeId, event.LastSize)
}
