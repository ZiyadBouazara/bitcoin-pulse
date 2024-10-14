package com.pricefetcherservice.domain;

public enum StockSymbols {
    BTC_USD;

    @Override
    public String  toString() {
        return name().replace("_", "-");
    }
}
