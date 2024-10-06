package com.pricefetcherservice.domain;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class Price {
    private double price;
    private double timestamp;

    @Override
    public String toString() {
        return "Price {" +
            "timestamp='" + timestamp + '\'' +
            ", price=" + price +
            '}';
    }

}
