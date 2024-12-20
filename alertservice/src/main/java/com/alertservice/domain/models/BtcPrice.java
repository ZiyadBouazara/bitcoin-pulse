package com.alertservice.domain.models;

import java.math.BigDecimal;

public class BtcPrice {

    private BigDecimal value;

    public BtcPrice(BigDecimal value) {
        this.value = value;
    }

    public BigDecimal getValue() {
        return value;
    }

    public void setValue(BigDecimal value) {
        this.value = value;
    }
}
