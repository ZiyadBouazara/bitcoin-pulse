package com.pricefetcherservice.infrastructure.kafka;


import com.pricefetcherservice.domain.PriceEvent;

public interface PriceProducer {
    void sendPrice(PriceEvent priceEvent);
}
