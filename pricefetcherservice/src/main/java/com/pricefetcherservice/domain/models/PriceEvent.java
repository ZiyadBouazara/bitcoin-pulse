package com.pricefetcherservice.domain.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;

@Data
@NoArgsConstructor
@Getter
@Setter
public class PriceEvent {
    @JsonProperty("type")
    private String type;

    @JsonProperty("sequence")
    private long sequence;

    @JsonProperty("product_id")
    private String productId;

    @JsonProperty("price")
    private String price;

    @JsonProperty("open_24h")
    private String open24h;

    @JsonProperty("volume_24h")
    private String volume24h;

    @JsonProperty("low_24h")
    private String low24h;

    @JsonProperty("high_24h")
    private String high24h;

    @JsonProperty("volume_30d")
    private String volume30d;

    @JsonProperty("best_bid")
    private String bestBid;

    @JsonProperty("best_bid_size")
    private String bestBidSize;

    @JsonProperty("best_ask")
    private String bestAsk;

    @JsonProperty("best_ask_size")
    private String bestAskSize;

    @JsonProperty("side")
    private String side;

    @JsonProperty("time")
    private String time;

    @JsonProperty("trade_id")
    private long tradeId;

    @JsonProperty("last_size")
    private String lastSize;
}
