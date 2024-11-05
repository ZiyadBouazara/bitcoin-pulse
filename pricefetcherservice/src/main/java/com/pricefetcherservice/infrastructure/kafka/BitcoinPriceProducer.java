package com.pricefetcherservice.infrastructure.kafka;

import com.pricefetcherservice.domain.models.PriceEvent;
import com.pricefetcherservice.domain.PriceProducer;
import org.apache.kafka.clients.admin.NewTopic;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class BitcoinPriceProducer implements PriceProducer {
    private static final Logger logger = LoggerFactory.getLogger(BitcoinPriceProducer.class);
    private final NewTopic bitcoinPriceTopic;
    private final KafkaTemplate<String, PriceEvent> kafkaTemplate;

    public BitcoinPriceProducer(NewTopic bitcoinPriceTopic, KafkaTemplate<String, PriceEvent> kafkaTemplate) {
        this.bitcoinPriceTopic = bitcoinPriceTopic;
        this.kafkaTemplate = kafkaTemplate;
    }

    @Override
    public void sendPrice(PriceEvent priceEvent) { // TODO: use protobuf instead of string
        logger.info("🚀 🚀 🚀 BTC Price Event sent 🚀 🚀 🚀 : \n{}", priceEvent.toString());

        kafkaTemplate.send(bitcoinPriceTopic.name(), priceEvent);
    }
}
