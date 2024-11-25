package dtos

import (
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"strconv"
	"time"
)

type PriceEventDTO struct {
	Type        string `json:"type"`
	Sequence    int64  `json:"sequence"`
	ProductID   string `json:"product_id"`
	Price       string `json:"price"`
	Open24H     string `json:"open_24h"`
	Volume24H   string `json:"volume_24h"`
	Low24H      string `json:"low_24h"`
	High24H     string `json:"high_24h"`
	Volume30D   string `json:"volume_30d"`
	BestBid     string `json:"best_bid"`
	BestBidSize string `json:"best_bid_size"`
	BestAsk     string `json:"best_ask"`
	BestAskSize string `json:"best_ask_size"`
	Side        string `json:"side"`
	Time        string `json:"time"`
	TradeId     int64  `json:"trade_id"`
	LastSize    string `json:"last_size"`
}

func (e *PriceEventDTO) FormatLog() string {
	return fmt.Sprintf(
		"\nType: %s\n"+
			"Sequence: %d\n"+
			"ProductID: %s\n"+
			"Price: %s\n"+
			"Open24H: %s\n"+
			"Volume24H: %s\n"+
			"Low24H: %s\n"+
			"High24H: %s\n"+
			"Volume30D: %s\n"+
			"BestBid: %s\n"+
			"BestBidSize: %s\n"+
			"BestAsk: %s\n"+
			"BestAskSize: %s\n"+
			"Side: %s\n"+
			"Time: %s\n"+
			"TradeID: %d\n"+
			"LastSize: %s",
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
		e.Time,
		e.TradeId,
		e.LastSize,
	)
}

func ToPriceEvent(dto *PriceEventDTO) (*domain.PriceEvent, error) {
	parseFloat := func(value string, fieldName string) (float64, error) {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, fmt.Errorf("error parsing %s: %v", fieldName, err)
		}
		return f, nil
	}

	if !domain.IsSupportedStock(dto.ProductID) {
		return nil, fmt.Errorf("unsupported stock: %v", dto.ProductID)
	}
	productID := domain.Stock(dto.ProductID)

	price, err := parseFloat(dto.Price, "Price")
	if err != nil {
		return nil, err
	}

	open24h, err := parseFloat(dto.Open24H, "Open24H")
	if err != nil {
		return nil, err
	}

	volume24h, err := parseFloat(dto.Volume24H, "Volume24H")
	if err != nil {
		return nil, err
	}

	low24h, err := parseFloat(dto.Low24H, "Low24H")
	if err != nil {
		return nil, err
	}

	high24h, err := parseFloat(dto.High24H, "High24H")
	if err != nil {
		return nil, err
	}

	volume30d, err := parseFloat(dto.Volume30D, "Volume30D")
	if err != nil {
		return nil, err
	}

	bestBid, err := parseFloat(dto.BestBid, "BestBid")
	if err != nil {
		return nil, err
	}

	bestBidSize, err := parseFloat(dto.BestBidSize, "BestBidSize")
	if err != nil {
		return nil, err
	}

	bestAsk, err := parseFloat(dto.BestAsk, "BestAsk")
	if err != nil {
		return nil, err
	}

	bestAskSize, err := parseFloat(dto.BestAskSize, "BestAskSize")
	if err != nil {
		return nil, err
	}

	lastSize, err := parseFloat(dto.LastSize, "LastSize")
	if err != nil {
		return nil, err
	}

	parsedTime, err := time.Parse(time.RFC3339, dto.Time)
	if err != nil {
		return nil, fmt.Errorf("error parsing Time: %v", err)
	}

	event := &domain.PriceEvent{
		Type:        dto.Type,
		Sequence:    dto.Sequence,
		ProductID:   productID,
		Price:       price,
		Open24H:     open24h,
		Volume24H:   volume24h,
		Low24H:      low24h,
		High24H:     high24h,
		Volume30D:   volume30d,
		BestBid:     bestBid,
		BestBidSize: bestBidSize,
		BestAsk:     bestAsk,
		BestAskSize: bestAskSize,
		Side:        dto.Side,
		Time:        parsedTime,
		TradeId:     dto.TradeId,
		LastSize:    lastSize,
	}

	return event, nil
}
