package com.pricefetcherservice.domain;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class PriceEvent {
    private String timestamp;
    private double price;
}
