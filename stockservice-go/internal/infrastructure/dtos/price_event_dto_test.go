package dtos

import (
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToPriceEvent_Success(t *testing.T) {
	dto := createValidDtoFixture()

	actualPriceEvent, err := ToPriceEvent(dto)
	expectedPriceEvent := expectedPriceEventFixture()

	assert.NoError(t, err)
	assert.Equal(t, expectedPriceEvent, actualPriceEvent)
}

func TestToPriceEvent_ErrorParsingFloat(t *testing.T) {
	dto := &PriceEventDTO{
		Price: "invalid",
	}

	event, err := ToPriceEvent(dto)

	assert.Nil(t, event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing Price")
}

func TestToPriceEvent_ErrorParsingTime(t *testing.T) {
	dto := &PriceEventDTO{
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
		Time:        "invalid",
		TradeId:     100,
		LastSize:    "100.0",
	}

	event, err := ToPriceEvent(dto)

	assert.Nil(t, event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing Time")
}

func createValidDtoFixture() *PriceEventDTO {
	return &PriceEventDTO{
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

func expectedPriceEventFixture() *domain.PriceEvent {
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
