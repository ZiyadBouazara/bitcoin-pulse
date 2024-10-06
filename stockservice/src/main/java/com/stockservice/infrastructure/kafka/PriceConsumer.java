package com.stockservice.infrastructure.kafka;

import com.stockservice.domain.PriceEvent;

public interface PriceConsumer {
    void consume(PriceEvent priceEvent);
}
