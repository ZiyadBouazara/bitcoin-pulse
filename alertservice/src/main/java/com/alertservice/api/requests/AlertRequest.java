package com.alertservice.api.requests;

import com.alertservice.domain.models.BtcPrice;

import java.math.BigDecimal;

public class AlertRequest {
    private String symbol;
    private String email;
    private String phone;
    private BtcPrice price;
    private BigDecimal priceDifference;
    private String text;

    public AlertRequest() {
    }

    public AlertRequest(String symbol, String email, String phone, BtcPrice price, BigDecimal priceDifference,
                        String text) {
        this.symbol = symbol;
        this.price = price;
        this.email = email;
        this.phone = phone;
        this.priceDifference = priceDifference;
        this.text = text;
    }

    public String getSymbol() {
        return symbol;
    }

    public void setSymbol(String symbol) {
        this.symbol = symbol;
    }

    public BtcPrice getPrice() {
        return price;
    }

    public void setPrice(BtcPrice price) {
        this.price = price;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public BigDecimal getPriceDifference() {
        return priceDifference;
    }

    public void setPriceDifference(BigDecimal priceDifference) {
        this.priceDifference = priceDifference;
    }
}
