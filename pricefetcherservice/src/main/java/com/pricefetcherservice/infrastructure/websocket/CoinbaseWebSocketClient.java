package com.pricefetcherservice.infrastructure.websocket;


import com.pricefetcherservice.domain.PriceUpdateListener;
import com.pricefetcherservice.domain.PriceWebSocketClient;
import com.pricefetcherservice.infrastructure.dtos.SubscribePriceDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.lang.NonNull;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.*;
import org.springframework.web.socket.client.standard.StandardWebSocketClient;

import java.util.List;

@Component
public class CoinbaseWebSocketClient implements PriceWebSocketClient {
    private static final Logger logger = LoggerFactory.getLogger(CoinbaseWebSocketClient.class);
    private static final List<String> CHANNELS = List.of("ticker_batch");
    private final String serverUri;
    private PriceUpdateListener priceUpdateListener;

    public CoinbaseWebSocketClient(@Value("${coinbase.ws.url}") String serverUri) {
        this.serverUri = serverUri;
        logger.info("WebSocket serverUri: " + this.serverUri);
    }

    @Override
    public void connect(List<String> stocks) {
        try {
            StandardWebSocketClient client = new StandardWebSocketClient();
            client.doHandshake(new CoinbaseWebSocketHandler(stocks), serverUri);
        } catch (Exception e) {
            logger.error("WebSocket connection error", e);
        }
    }

    @Override
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
        public void handleMessage(@NonNull WebSocketSession session, WebSocketMessage<?> message) {
            String payload = message.getPayload().toString();
            logger.debug("📩 📩 📩 Received message from Coinbase Websocket 📩 📩 📩 : \n" + payload);
            if (priceUpdateListener != null) {
                priceUpdateListener.onPriceUpdate(payload);
            }
        }

        @Override
        public void handleTransportError(@NonNull WebSocketSession session, @NonNull Throwable exception) {
            logger.error("WebSocket transport error", exception);
        }

        @Override
        public void afterConnectionClosed(@NonNull WebSocketSession session, @NonNull CloseStatus closeStatus) {
            logger.info("WebSocket connection closed: " + closeStatus);
        }

        @Override
        public boolean supportsPartialMessages() {
            return false;
        }
    }
}

