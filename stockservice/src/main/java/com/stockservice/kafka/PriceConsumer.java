package com.stockservice.kafka;

import com.basedomain.dto.PriceEvent;

public interface PriceConsumer {
    void consume(PriceEvent priceEvent);
}
