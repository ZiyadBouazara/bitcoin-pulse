package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;
import com.alertservice.domain.models.TrendDirection;

import java.math.BigDecimal;

public class AlertFactory {
    public Alert createPriceAboveAlert(BtcPrice bitcoinPrice, BigDecimal priceDifference) {
        return new Alert(
                "Price over threshold",
                "The price of Bitcoin has gone over the threshold by: " + priceDifference,
                "email@email.com",
                "null",
                bitcoinPrice,
                TrendDirection.UP
        );
    }

    public Alert createPriceBelowAlert(BtcPrice bitcoinPrice, BigDecimal priceDifference) {
        return new Alert(
                "Price below threshold",
                "The price of Bitcoin has passed under the threshold by: " + priceDifference,
                "email@email.com",
                "null",
                bitcoinPrice,
                TrendDirection.UP
        );
    }
}
