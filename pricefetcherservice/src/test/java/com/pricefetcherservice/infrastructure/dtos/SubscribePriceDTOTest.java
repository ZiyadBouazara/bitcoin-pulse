package com.pricefetcherservice.infrastructure.dtos;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;


class SubscribePriceDTOTest {
    private static final List<String> CHANNELS = List.of("ticker", "heartbeat");
    private static final List<String> STOCKS = List.of("BTC-USD", "ETH-USD");
    private SubscribePriceDTO dto;

    @BeforeEach
    public void setUp() {
        dto = new SubscribePriceDTO(CHANNELS, STOCKS);
    }

    @Test
    public void givenValidDTO_whenToJson_thenReturnValidJson() {
        String jsonString = dto.toJson();

        assertNotNull(jsonString);
        assertTrue(jsonString.contains("\"type\":\"subscribe\""));
        assertTrue(jsonString.contains("\"channels\":[\"ticker\",\"heartbeat\"]"));
        assertTrue(jsonString.contains("\"product_ids\":[\"BTC-USD\",\"ETH-USD\"]"));
    }
}