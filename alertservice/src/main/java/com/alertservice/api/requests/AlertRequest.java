package com.alertservice.api.requests;

import com.alertservice.domain.models.BtcPrice;

public class AlertRequest {
    private String symbol;
    private String email;
    private String phone;
    private BtcPrice price;
    private String text;

    public AlertRequest() {
    }

    public AlertRequest(String symbol, String email, String phone, BtcPrice price, String text) {
        this.symbol = symbol;
        this.price = price;
        this.email = email;
        this.phone = phone;
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
}
