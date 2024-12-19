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
    private final AlertRepository alertRepository;
    private final AlertFactory alertFactory;
    private final AlertSender alertSender;

    public PriceFilter(AlertRepository alertRepository,
                       AlertSender alertSender,
                       AlertFactory alertFactory) {
        this.alertRepository = alertRepository;
        this.alertSender = alertSender;
        this.alertFactory = alertFactory;
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
            alertRepository.saveAlert(
                    alertFactory.createPriceAboveAlert(currentPrice, priceDifference));
        } else {
            // todo: alertSend.sendAlert(alertFactory.createPriceBelowAlert(currentPrice, priceDifference));
            alertRepository.saveAlert(
                    alertFactory.createPriceBelowAlert(currentPrice, priceDifference));
        }

        comparisonPrice = currentPrice;
    }
}
