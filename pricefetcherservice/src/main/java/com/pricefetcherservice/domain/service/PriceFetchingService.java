package com.pricefetcherservice.domain.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.pricefetcherservice.domain.PriceEvent;
import com.pricefetcherservice.domain.PriceProducer;
import com.pricefetcherservice.domain.StockSymbols;
import com.pricefetcherservice.infrastructure.websocket.CoinbaseWebSocketClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;

import java.util.ArrayList;
import java.util.List;

@Service
public class PriceFetchingService {
    private static final Logger logger = LoggerFactory.getLogger(PriceFetchingService.class);
    private static final List<StockSymbols> SUPPORTED_STOCKS = new ArrayList<>(List.of(StockSymbols.BTC_USD));
    private PriceProducer bitcoinPriceProducer;
    private CoinbaseWebSocketClient coinbaseWebSocketClient;

    @EventListener(ApplicationReadyEvent.class)
    public void startFetching() {
        coinbaseWebSocketClient.connect(SUPPORTED_STOCKS.stream().map(StockSymbols::toString).toList());
        coinbaseWebSocketClient.setPriceUpdateListener(this::handlePriceUpdate);
    }

    public PriceFetchingService(PriceProducer bitcoinPriceProducer, CoinbaseWebSocketClient coinbaseWebSocketClient) {
        this.bitcoinPriceProducer = bitcoinPriceProducer;
        this.coinbaseWebSocketClient = coinbaseWebSocketClient;
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
