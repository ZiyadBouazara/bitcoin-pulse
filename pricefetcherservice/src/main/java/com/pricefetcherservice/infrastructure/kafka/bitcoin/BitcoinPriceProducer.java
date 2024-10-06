package com.pricefetcherservice.infrastructure.kafka.bitcoin;

import com.pricefetcherservice.domain.PriceEvent;
import com.pricefetcherservice.infrastructure.kafka.PriceProducer;
import org.apache.kafka.clients.admin.NewTopic;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.Message;
import org.springframework.messaging.support.MessageBuilder;
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
    public void sendPrice(PriceEvent priceEvent) {
        logger.info(String.format("Bitcoin Price Event => %s", priceEvent));

        Message<PriceEvent> message =
            MessageBuilder.withPayload(priceEvent).setHeader(KafkaHeaders.TOPIC, bitcoinPriceTopic.name()).build();

        kafkaTemplate.send(message);
    }
}
