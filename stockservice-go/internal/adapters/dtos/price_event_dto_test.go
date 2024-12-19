package dtos_test

import (
	"testing"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/dtos"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestToPriceEvent_Success(t *testing.T) {
	dto := testutils.CreateValidPriceEventDTO()

	actualPriceEvent, err := dtos.ToPriceEvent(dto)
	expectedPriceEvent := testutils.CreateValidPriceEvent()

	assert.NoError(t, err)
	assert.Equal(t, expectedPriceEvent, actualPriceEvent)
}

func TestToPriceEvent_ErrorParsingFloat(t *testing.T) {
	dto := &dtos.PriceEventDTO{
		ProductID: "BTC-USD",
		Price:     "invalid",
	}

	event, err := dtos.ToPriceEvent(dto)

	assert.Nil(t, event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing Price")
}

func TestToPriceEvent_ErrorParsingTime(t *testing.T) {
	dto := &dtos.PriceEventDTO{
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

	event, err := dtos.ToPriceEvent(dto)

	assert.Nil(t, event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing Time")
}
