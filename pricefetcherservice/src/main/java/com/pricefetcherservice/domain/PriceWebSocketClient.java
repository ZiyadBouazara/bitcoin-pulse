package com.pricefetcherservice.domain;

import java.util.List;

public interface PriceWebSocketClient {
    void connect(List<String> stocks);
    void setPriceUpdateListener(PriceUpdateListener listener);
}
