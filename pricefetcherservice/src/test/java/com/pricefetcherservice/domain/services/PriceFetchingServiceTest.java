package com.pricefetcherservice.domain.services;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.pricefetcherservice.domain.PriceProducer;
import com.pricefetcherservice.domain.PriceUpdateListener;
import com.pricefetcherservice.domain.PriceWebSocketClient;
import com.pricefetcherservice.domain.fixtures.PriceEventFixture;
import com.pricefetcherservice.domain.models.PriceEvent;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Captor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.doThrow;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;

class PriceFetchingServiceTest {
    private static final List<String> SUPPORTED_STOCKS = List.of("BTC-USD");
    private static final String AN_INVALID_JSON = "Invalid JSON";

    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Mock
    private PriceProducer bitcoinPriceProducer;

    @Mock
    private PriceWebSocketClient priceWebSocketClient;

    @InjectMocks
    private PriceFetchingService priceFetchingService;

    @Captor
    private ArgumentCaptor<PriceUpdateListener> priceUpdateListenerCaptor;

    @BeforeEach
    public void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    public void givenApplicationReadyEvent_whenStartFetching_thenConnectsWithCorrectStockSymbols() {
        priceFetchingService.startFetching();

        verify(priceWebSocketClient).setPriceUpdateListener(any());
        verify(priceWebSocketClient).connect(SUPPORTED_STOCKS);
    }

    @Test
    public void givenExceptionInStartFetching_whenStartFetching_thenThrowException() throws RuntimeException {
        doThrow(new RuntimeException()).when(priceWebSocketClient).connect(any());

        assertThrows(RuntimeException.class, () -> priceFetchingService.startFetching());
    }

    @Test
    public void givenValidPriceUpdate_whenHandlePriceUpdate_thenSendsPriceToProducer() throws JsonProcessingException {
        String validPriceUpdate = PriceEventFixture.builder().buildJson();

        PriceEvent expectedPriceEvent = objectMapper.readValue(validPriceUpdate, PriceEvent.class);

        priceFetchingService.startFetching();
        verify(priceWebSocketClient).setPriceUpdateListener(priceUpdateListenerCaptor.capture());
        PriceUpdateListener priceUpdateListener = priceUpdateListenerCaptor.getValue();

        priceUpdateListener.onPriceUpdate(validPriceUpdate);

        verify(bitcoinPriceProducer).sendPrice(expectedPriceEvent);
    }

    @Test
    public void givenInvalidPriceUpdate_whenHandlePriceUpdate_thenDoesNotSendPrice() {
        priceFetchingService.startFetching();
        verify(priceWebSocketClient).setPriceUpdateListener(priceUpdateListenerCaptor.capture());
        PriceUpdateListener priceUpdateListener = priceUpdateListenerCaptor.getValue();

        priceUpdateListener.onPriceUpdate(AN_INVALID_JSON);

        verify(bitcoinPriceProducer, never()).sendPrice(any());
    }
}