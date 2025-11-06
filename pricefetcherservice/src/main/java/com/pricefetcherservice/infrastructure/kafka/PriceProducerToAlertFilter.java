package com.pricefetcherservice.infrastructure.kafka;

import com.pricefetcherservice.domain.PriceProducer;
import com.pricefetcherservice.domain.models.PriceEvent;

public class PriceProducerToAlertFilter implements PriceProducer {
    @Override
    public void sendPrice(PriceEvent priceEvent) {

    }
}
