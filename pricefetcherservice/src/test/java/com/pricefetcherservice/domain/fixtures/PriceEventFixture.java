package com.pricefetcherservice.domain.fixtures;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.pricefetcherservice.domain.models.PriceEvent;

public class PriceEventFixture {
    private String type = "ticker";
    private long sequence = 1234567890L;
    private String productId = "BTC-USD";
    private String price = "50000.0";
    private String open24h = "48000.0";
    private String volume24h = "1000.0";
    private String low24h = "47000.0";
    private String high24h = "51000.0";
    private String volume30d = "30000.0";
    private String bestBid = "49900.0";
    private String bestBidSize = "1.5";
    private String bestAsk = "50100.0";
    private String bestAskSize = "2.0";
    private String side = "buy";
    private String time = "2024-10-14T15:38:52.418-04:00";
    private long tradeId = 987654321L;
    private String lastSize = "0.5";

    private static final ObjectMapper objectMapper = new ObjectMapper();

    public static PriceEventFixture builder() {
        return new PriceEventFixture();
    }

    public PriceEventFixture withType(String type) {
        this.type = type;
        return this;
    }

    public PriceEventFixture withSequence(long sequence) {
        this.sequence = sequence;
        return this;
    }

    public PriceEventFixture withProductId(String productId) {
        this.productId = productId;
        return this;
    }

    public PriceEventFixture withPrice(String price) {
        this.price = price;
        return this;
    }

    public PriceEventFixture withOpen24h(String open24h) {
        this.open24h = open24h;
        return this;
    }

    public PriceEventFixture withVolume24h(String volume24h) {
        this.volume24h = volume24h;
        return this;
    }

    public PriceEventFixture withLow24h(String low24h) {
        this.low24h = low24h;
        return this;
    }

    public PriceEventFixture withHigh24h(String high24h) {
        this.high24h = high24h;
        return this;
    }

    public PriceEventFixture withVolume30d(String volume30d) {
        this.volume30d = volume30d;
        return this;
    }

    public PriceEventFixture withBestBid(String bestBid) {
        this.bestBid = bestBid;
        return this;
    }

    public PriceEventFixture withBestBidSize(String bestBidSize) {
        this.bestBidSize = bestBidSize;
        return this;
    }

    public PriceEventFixture withBestAsk(String bestAsk) {
        this.bestAsk = bestAsk;
        return this;
    }

    public PriceEventFixture withBestAskSize(String bestAskSize) {
        this.bestAskSize = bestAskSize;
        return this;
    }

    public PriceEventFixture withSide(String side) {
        this.side = side;
        return this;
    }

    public PriceEventFixture withTime(String time) {
        this.time = time;
        return this;
    }

    public PriceEventFixture withTradeId(long tradeId) {
        this.tradeId = tradeId;
        return this;
    }

    public PriceEventFixture withLastSize(String lastSize) {
        this.lastSize = lastSize;
        return this;
    }

    public PriceEvent build() {
        return new PriceEvent(this.type, this.sequence, this.productId, this.price, this.open24h, this.volume24h, this.low24h,
            this.high24h, this.volume30d, this.bestBid, this.bestBidSize, this.bestAsk, this.bestAskSize, this.side, this.time,
            this.tradeId, this.lastSize);
    }

    public String buildJson() throws JsonProcessingException {
        PriceEvent priceEvent = build();
        return objectMapper.writeValueAsString(priceEvent);
    }
}
