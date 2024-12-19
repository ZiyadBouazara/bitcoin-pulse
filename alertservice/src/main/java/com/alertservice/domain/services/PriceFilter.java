package com.alertservice.domain.services;

import com.alertservice.domain.models.BtcPrice;
import com.pricefetcherservice.domain.models.PriceEvent;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;

@Service
public class PriceFilter {

    private final BtcPrice BTC_PRICE_MARGIN = new BtcPrice(new BigDecimal(100));
    private BtcPrice comparisonPrice = null; // Initially null to signify no price yet
    private final AlertService alertService;

    public PriceFilter(AlertService alertService) {
        this.alertService = alertService;
    }

    @KafkaListener(topics = "bitcoin-price-topic", groupId = "price-filter-group", containerFactory = "kafkaListenerContainerFactory")
    public void processPriceEvent(PriceEvent priceEvent) {
        BigDecimal currentPriceValue = new BigDecimal(priceEvent.price());
        BtcPrice currentPrice = new BtcPrice(currentPriceValue);

        if (comparisonPrice == null) {
            comparisonPrice = currentPrice;
            System.out.println("Initialized COMPARISON_PRICE: " + comparisonPrice.getValue());
            return;
        }

        BigDecimal priceDifference = currentPrice.getValue().subtract(comparisonPrice.getValue()).abs();

        if (priceDifference.compareTo(BTC_PRICE_MARGIN.getValue()) > 0) {
            // define which users to send the alert to then:
            // todo: alertSend.sendAlert(alertFactory.createPriceAboveAlert(currentPrice, priceDifference));
            alertService.sendTriggeredAlerts(priceEvent.price());

        } else {
            // todo: alertSend.sendAlert(alertFactory.createPriceBelowAlert(currentPrice, priceDifference));
            alertService.sendTriggeredAlerts(priceEvent.price());
        }

        comparisonPrice = currentPrice;
    }
}
