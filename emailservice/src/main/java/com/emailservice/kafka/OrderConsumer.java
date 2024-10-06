package com.emailservice.kafka;

import com.basedomain.dto.OrderEvent;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class OrderConsumer {
    public static final Logger logger = LoggerFactory.getLogger(OrderConsumer.class);


    @KafkaListener(topics = "${spring.kafka.topic.name}",groupId = "${spring.kafka.consumer.group-id}")
    public void consume(OrderEvent orderEvent) {
        logger.info("Order event recieved in service => %s",orderEvent.toString());
    }


}
