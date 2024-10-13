package com.pricefetcherservice.infrastructure;

import com.pricefetcherservice.domain.Price;
import com.pricefetcherservice.domain.StockSymbols;
import com.pricefetcherservice.infrastructure.dtos.CoinbasePriceResponse;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

@Component
public class CoinbaseApiClient {
    @Value("${coinbase.api.url}")
    private String coinbaseApiUrl;
    private final RestTemplate restTemplate;

    public CoinbaseApiClient(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    public Price getPrice(StockSymbols stockSymbol) {
        CoinbasePriceResponse response =
            restTemplate.getForObject(coinbaseApiUrl, CoinbasePriceResponse.class); // TODO: fetch only for for param stocksymbol

//        if (response != null && response.getData() != null) {
//            return new Price(stockSymbol, response.getData().getAmount());
//        }
        return null;
    }
}
