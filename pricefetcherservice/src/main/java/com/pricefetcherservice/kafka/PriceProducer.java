package com.pricefetcherservice.kafka;

import com.basedomain.dto.PriceEvent;

public interface PriceProducer {
    void sendPrice(PriceEvent priceEvent);
}
