package com.pricefetcherservice.domain.models;


import com.fasterxml.jackson.annotation.JsonProperty;

public record PriceEvent(
    @JsonProperty("type")
    String type,

    @JsonProperty("sequence")
    long sequence,

    @JsonProperty("product_id")
    String productId,

    @JsonProperty("price")
    String price,

    @JsonProperty("open_24h")
    String open24h,

    @JsonProperty("volume_24h")
    String volume24h,

    @JsonProperty("low_24h")
    String low24h,

    @JsonProperty("high_24h")
    String high24h,

    @JsonProperty("volume_30d")
    String volume30d,

    @JsonProperty("best_bid")
    String bestBid,

    @JsonProperty("best_bid_size")
    String bestBidSize,

    @JsonProperty("best_ask")
    String bestAsk,

    @JsonProperty("best_ask_size")
    String bestAskSize,

    @JsonProperty("side")
    String side,

    @JsonProperty("time")
    String time,

    @JsonProperty("trade_id")
    long tradeId,

    @JsonProperty("last_size")
    String lastSize
) {
}
