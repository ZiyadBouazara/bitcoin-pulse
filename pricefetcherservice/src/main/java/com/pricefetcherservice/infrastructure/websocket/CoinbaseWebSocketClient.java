package com.pricefetcherservice.infrastructure.websocket;

import com.neovisionaries.ws.client.WebSocket;
import com.neovisionaries.ws.client.WebSocketAdapter;
import com.neovisionaries.ws.client.WebSocketException;
import com.neovisionaries.ws.client.WebSocketFactory;
import com.neovisionaries.ws.client.WebSocketState;
import com.pricefetcherservice.infrastructure.dtos.SubscribePriceDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import java.io.IOException;
import java.util.List;

@Component
public class CoinbaseWebSocketClient {
    private static final Logger logger = LoggerFactory.getLogger(CoinbaseWebSocketClient.class);
    private String serverUri;
    private WebSocket ws;
    private PriceUpdateListener priceUpdateListener;

    public CoinbaseWebSocketClient(@Value("${coinbase.ws.url}") String serverUri) {
        this.serverUri = serverUri;
    }

    public void connect(List<String> stocks) {
        try {
            ws = new WebSocketFactory().createSocket(serverUri);

            ws.addListener(new WebSocketAdapter() {
                @Override
                public void onStateChanged(WebSocket websocket, WebSocketState newState) {
                    if (newState == WebSocketState.OPEN) {
                        String subscribeMessage = new SubscribePriceDTO(stocks).toString();
                        ws.sendText(subscribeMessage);
                    }
                }

                @Override
                public void onTextMessage(WebSocket websocket, String message) {
                    if (priceUpdateListener != null) {
                        priceUpdateListener.onPriceUpdate(message);
                    }
                    logger.info("Received message: " + message);
                }

                @Override
                public void onError(WebSocket websocket, WebSocketException cause) {
                    logger.error("WebSocket Error: " + cause.getMessage());
                }
            });

            ws.connect();

        } catch (IOException | WebSocketException e) {
            logger.error("WebSocket Error: " + e.getMessage());
        }
    }

    public void setPriceUpdateListener(PriceUpdateListener listener) {
        this.priceUpdateListener = listener;
    }
}
