package com.pricefetcherservice.domain.models;

public enum StockSymbols {
    BTC_USD;

    @Override
    public String  toString() {
        return name().replace("_", "-");
    }
}
