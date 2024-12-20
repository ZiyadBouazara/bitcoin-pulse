package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;
import com.alertservice.domain.models.TrendDirection;

import java.math.BigDecimal;

public class AlertFactory {

    public Alert createPriceAlert(BtcPrice bitcoinPrice, BigDecimal priceDifference) {
        return new Alert(
                "Price Alert",
                "The price of Bitcoin has changed by: " + priceDifference,
                "email@email.com",
                "phone-number",
                bitcoinPrice
        );
    }

    public Alert createPriceAboveAlert(BtcPrice bitcoinPrice, BigDecimal priceDifference) {
        return new Alert(
                "Price over threshold",
                "The price of Bitcoin has gone over the threshold by: " + priceDifference,
                "email@email.com",
                "null",
                bitcoinPrice
        );
    }

    public Alert createPriceBelowAlert(BtcPrice bitcoinPrice, BigDecimal priceDifference) {
        return new Alert(
                "Price below threshold",
                "The price of Bitcoin has passed under the threshold by: " + priceDifference,
                "email@email.com",
                "null",
                bitcoinPrice
        );
    }
}
