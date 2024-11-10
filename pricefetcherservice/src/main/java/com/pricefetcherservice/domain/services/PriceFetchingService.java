package com.pricefetcherservice.domain.services;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.pricefetcherservice.domain.models.PriceEvent;
import com.pricefetcherservice.domain.PriceProducer;
import com.pricefetcherservice.domain.PriceWebSocketClient;
import com.pricefetcherservice.domain.models.StockSymbols;
import com.pricefetcherservice.infrastructure.websocket.CoinbaseWebSocketClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;

import java.util.ArrayList;
import java.util.List;

@Service
public class PriceFetchingService {
    private static final Logger logger = LoggerFactory.getLogger(PriceFetchingService.class);
    private static final List<StockSymbols> SUPPORTED_STOCKS = new ArrayList<>(List.of(StockSymbols.BTC_USD));
    private final PriceProducer bitcoinPriceProducer;
    private final PriceWebSocketClient coinbaseWebSocketClient;

    public PriceFetchingService(PriceProducer bitcoinPriceProducer, CoinbaseWebSocketClient coinbaseWebSocketClient) {
        this.bitcoinPriceProducer = bitcoinPriceProducer;
        this.coinbaseWebSocketClient = coinbaseWebSocketClient;
    }

    @Async
    @EventListener(ApplicationReadyEvent.class)
    public void startFetching() {
        try {
            logger.info("PriceFetchingService has started fetching bitcoin prices");
            coinbaseWebSocketClient.setPriceUpdateListener(this::handlePriceUpdate);
            coinbaseWebSocketClient.connect(SUPPORTED_STOCKS.stream().map(StockSymbols::toString).toList());
        } catch (Exception e) {
            logger.error("Exception in startFetching(): " + e.getMessage());
        }
    }

    private void handlePriceUpdate(String priceUpdate) {
        try {
            PriceEvent priceEvent = new ObjectMapper().readValue(priceUpdate, PriceEvent.class);
            bitcoinPriceProducer.sendPrice(priceEvent);
        } catch (JsonProcessingException e) {
            logger.error("Json Parsing Error: " + e.getMessage());
        }
    }
}