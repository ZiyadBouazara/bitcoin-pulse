package com.pricefetcherservice.infrastructure.websocket;


import com.pricefetcherservice.infrastructure.dtos.SubscribePriceDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.*;
import org.springframework.web.socket.client.standard.StandardWebSocketClient;

import java.util.List;

@Component
public class CoinbaseWebSocketClient {
    private static final Logger logger = LoggerFactory.getLogger(CoinbaseWebSocketClient.class);
    private static final List<String> CHANNELS = List.of("ticker_batch");
    private final String serverUri;
    private PriceUpdateListener priceUpdateListener;

    public CoinbaseWebSocketClient(@Value("${coinbase.ws.url}") String serverUri) {
        this.serverUri = serverUri;
        logger.info("WebSocket serverUri: " + this.serverUri);
    }

    public void connect(List<String> stocks) {
        try {
            StandardWebSocketClient client = new StandardWebSocketClient();
            client.doHandshake(new CoinbaseWebSocketHandler(stocks), serverUri);
        } catch (Exception e) {
            logger.error("WebSocket connection error", e);
        }
    }

    public void setPriceUpdateListener(PriceUpdateListener listener) {
        this.priceUpdateListener = listener;
    }

    private class CoinbaseWebSocketHandler implements WebSocketHandler {
        private final List<String> stocks;

        public CoinbaseWebSocketHandler(List<String> stocks) {
            this.stocks = stocks;
        }

        @Override
        public void afterConnectionEstablished(WebSocketSession session) throws Exception {
            logger.info("WebSocket connection established");
            String subscribeMessage = new SubscribePriceDTO(CHANNELS, stocks).toString();
            session.sendMessage(new TextMessage(subscribeMessage));
            logger.info("Subscription message sent: " + subscribeMessage);
        }

        @Override
        public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) throws Exception {
            String payload = message.getPayload().toString();
            logger.info("Received message: " + payload);
            if (priceUpdateListener != null) {
                priceUpdateListener.onPriceUpdate(payload);
            }
        }

        @Override
        public void handleTransportError(WebSocketSession session, Throwable exception) throws Exception {
            logger.error("WebSocket transport error", exception);
        }

        @Override
        public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus) throws Exception {
            logger.info("WebSocket connection closed: " + closeStatus);
        }

        @Override
        public boolean supportsPartialMessages() {
            return false;
        }
    }
}

