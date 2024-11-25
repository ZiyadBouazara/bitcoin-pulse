package testutils

import (
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/dtos"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"time"
)

func CreateValidPriceEvent() *domain.PriceEvent {
	expectedTime, _ := time.Parse(time.RFC3339, "2023-11-18T12:34:56Z")

	return &domain.PriceEvent{
		Type:        "ticker",
		Sequence:    100,
		ProductID:   "BTC-USD",
		Price:       100.0,
		Open24H:     100.0,
		Volume24H:   100.0,
		Low24H:      100.0,
		High24H:     100.0,
		Volume30D:   100.0,
		BestBid:     100.0,
		BestBidSize: 100.0,
		BestAsk:     100.0,
		BestAskSize: 100.0,
		Side:        "buy",
		Time:        expectedTime,
		TradeId:     100.0,
		LastSize:    100.0,
	}
}

func CreateValidPriceEventDTO() *dtos.PriceEventDTO {
	return &dtos.PriceEventDTO{
		Type:        "ticker",
		Sequence:    100,
		ProductID:   "BTC-USD",
		Price:       "100.0",
		Open24H:     "100.0",
		Volume24H:   "100.0",
		Low24H:      "100.0",
		High24H:     "100.0",
		Volume30D:   "100.0",
		BestBid:     "100.0",
		BestBidSize: "100.0",
		BestAsk:     "100.0",
		BestAskSize: "100.0",
		Side:        "buy",
		Time:        "2023-11-18T12:34:56Z",
		TradeId:     100,
		LastSize:    "100.0",
	}
}
