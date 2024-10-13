package com.pricefetcherservice.domain;


public interface PriceProducer {
    void sendPrice(PriceEvent priceEvent);
}
