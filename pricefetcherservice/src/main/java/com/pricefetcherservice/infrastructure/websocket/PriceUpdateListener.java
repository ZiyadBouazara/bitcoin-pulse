package com.pricefetcherservice.infrastructure.websocket;

public interface PriceUpdateListener {
    void onPriceUpdate(String message);
}
