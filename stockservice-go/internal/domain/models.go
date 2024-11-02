package domain

type PriceEvent struct {
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
