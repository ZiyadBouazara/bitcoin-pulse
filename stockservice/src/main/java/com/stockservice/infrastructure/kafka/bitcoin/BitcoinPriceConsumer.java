package com.stockservice.infrastructure.kafka.bitcoin;

import com.stockservice.domain.PriceEvent;
import com.stockservice.infrastructure.kafka.PriceConsumer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class BitcoinPriceConsumer implements PriceConsumer {
    public static final Logger logger = LoggerFactory.getLogger(BitcoinPriceConsumer.class);

    @KafkaListener(topics = "${spring.kafka.topic.name.bitcoin.price}", groupId = "${spring.kafka.consumer.group-id}")
    public void consume(PriceEvent priceEvent) { // TODO: create PriceEventModel, then map it to PriceEvent domain object
        logger.info("Bitcoin Price Event received in stock service => {}", priceEvent.toString());
    }
}
