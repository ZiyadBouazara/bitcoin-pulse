package com.pricefetcherservice.domain;


import com.pricefetcherservice.domain.models.PriceEvent;

public interface PriceProducer {
    void sendPrice(PriceEvent priceEvent);
}
