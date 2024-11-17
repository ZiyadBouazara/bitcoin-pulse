package com.pricefetcherservice.infrastructure.websocket;


import com.pricefetcherservice.domain.PriceUpdateListener;
import com.pricefetcherservice.infrastructure.dtos.SubscribePriceDTO;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;

import java.util.List;

import static org.mockito.Mockito.verify;


import static org.junit.jupiter.api.Assertions.assertEquals;

public class CoinbaseWebSocketClientTest {
    private static final String SERVER_URI = "wss://example.com";
    private static final List<String> STOCKS = List.of("BTC-USD", "ETH-USD");

    @Mock
    private PriceUpdateListener priceUpdateListener;

    @Mock
    private WebSocketSession session;

    @Mock
    private CoinbaseWebSocketClient.CoinbaseWebSocketHandler handler;

    @BeforeEach
    public void setUp() {
        MockitoAnnotations.openMocks(this);
        CoinbaseWebSocketClient client = new CoinbaseWebSocketClient(SERVER_URI);
        client.setPriceUpdateListener(priceUpdateListener);
        handler = client.new CoinbaseWebSocketHandler(STOCKS);
    }

    @Test
    public void givenStocks_whenAfterConnectionEstablished_thenSendsSubscribeMessage() throws Exception {
        handler.afterConnectionEstablished(session);

        ArgumentCaptor<TextMessage> captor = ArgumentCaptor.forClass(TextMessage.class);
        verify(session).sendMessage(captor.capture());
        String sentMessage = captor.getValue().getPayload();

        String expectedMessage = new SubscribePriceDTO(CoinbaseWebSocketClient.CHANNELS, STOCKS).toJson();

        assertEquals(expectedMessage, sentMessage);
    }

    @Test
    public void givenMessage_whenHandleMessage_thenPriceUpdateListenerCalled() {
        String payload = "{\"type\":\"ticker_batch\",\"data\":[{\"symbol\":\"BTC-USD\",\"price\":50000}]}";
        TextMessage message = new TextMessage(payload);

        handler.handleMessage(session, message);

        verify(priceUpdateListener).onPriceUpdate(payload);
    }
}
