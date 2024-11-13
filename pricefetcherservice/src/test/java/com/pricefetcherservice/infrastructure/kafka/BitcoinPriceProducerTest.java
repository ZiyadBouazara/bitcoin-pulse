package com.pricefetcherservice.infrastructure.kafka;


import com.pricefetcherservice.domain.fixtures.PriceEventFixture;
import com.pricefetcherservice.domain.models.PriceEvent;
import org.apache.kafka.clients.admin.NewTopic;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.kafka.core.KafkaTemplate;

import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

class BitcoinPriceProducerTest {
    private static final String A_TOPIC_NAME = "a name";

    @Mock
    private NewTopic bitcoinPriceTopic;

    @Mock
    private KafkaTemplate<String, PriceEvent> kafkaTemplate;

    @InjectMocks
    private BitcoinPriceProducer bitcoinPriceProducer;

    @BeforeEach
    public void setUp() {
        MockitoAnnotations.openMocks(this);

        when(bitcoinPriceTopic.name()).thenReturn(A_TOPIC_NAME);
    }

    @Test
    public void givenValidPriceEvent_whenSendPrice_thenInvokeBrokerSend() {
        PriceEvent priceEvent = PriceEventFixture.builder().build();

        bitcoinPriceProducer.sendPrice(priceEvent);

        verify(kafkaTemplate).send(A_TOPIC_NAME, priceEvent);
    }
}