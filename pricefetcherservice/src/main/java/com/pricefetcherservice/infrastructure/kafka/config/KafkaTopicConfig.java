package com.pricefetcherservice.infrastructure.kafka.config;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
public class KafkaTopicConfig {
    @Value("${spring.kafka.bitcoin.price.topic.name}")
    private String bitcoinPriceTopicName;

    @Bean
    public NewTopic bitcoinPriceTopic() {
        return TopicBuilder.name(bitcoinPriceTopicName).build();
    }
}
