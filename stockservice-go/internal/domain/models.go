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

func (e *PriceEvent) FormatLog() string {
	return fmt.Sprintf(
		"\nType: %s\n"+
			"Sequence: %d\n"+
			"ProductID: %s\n"+
			"Price: %.2f\n"+
			"Open24H: %.2f\n"+
			"Volume24H: %.2f\n"+
			"Low24H: %.2f\n"+
			"High24H: %.2f\n"+
			"Volume30D: %.2f\n"+
			"BestBid: %.2f\n"+
			"BestBidSize: %.2f\n"+
			"BestAsk: %.2f\n"+
			"BestAskSize: %.2f\n"+
			"Side: %s\n"+
			"Time: %s\n"+
			"TradeID: %d\n"+
			"LastSize: %.2f",
		e.Type,
		e.Sequence,
		e.ProductID,
		e.Price,
		e.Open24H,
		e.Volume24H,
		e.Low24H,
		e.High24H,
		e.Volume30D,
		e.BestBid,
		e.BestBidSize,
		e.BestAsk,
		e.BestAskSize,
		e.Side,
		e.Time.Format(time.RFC3339),
		e.TradeId,
		e.LastSize,
	)
}
