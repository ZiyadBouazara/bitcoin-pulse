package com.pricefetcherservice.domain.service;

import com.pricefetcherservice.domain.Price;
import com.pricefetcherservice.domain.PriceEvent;
import com.pricefetcherservice.domain.PriceProducer;
import com.pricefetcherservice.domain.StockSymbols;
import com.pricefetcherservice.infrastructure.CoinbaseApiClient;
import org.springframework.stereotype.Service;

@Service
public class PriceFetchingService {
    private PriceProducer bitcoinPriceProducer;
    private CoinbaseApiClient coinbaseApiClient;

    public PriceFetchingService(PriceProducer bitcoinPriceProducer, CoinbaseApiClient coinbaseApiClient) {
        this.bitcoinPriceProducer = bitcoinPriceProducer;
        this.coinbaseApiClient = coinbaseApiClient;
    }

    public void fetchAndPublishBitcoinPrice() {
        Price bitcoinPrice = coinbaseApiClient.getPrice(StockSymbols.BTC);
        PriceEvent priceEvent = new PriceEvent(bitcoinPrice.getTimestamp(), bitcoinPrice.getPrice());
        bitcoinPriceProducer.sendPrice(priceEvent);
    }
}
