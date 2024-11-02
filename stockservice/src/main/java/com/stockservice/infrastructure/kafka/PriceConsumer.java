package com.stockservice.infrastructure.kafka;

import com.stockservice.infrastructure.models.PriceEventModel;

public interface PriceConsumer {
    void consume(PriceEventModel priceEventModel);
}
