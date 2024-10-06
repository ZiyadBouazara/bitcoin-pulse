package com.basedomain.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class PriceEvent {
    private String message;
    private String status;
    private Price price;

    @Override
    public String toString() {
        return "PriceEvent {" +
            "message='" + message + '\'' +
            ", status='" + status + '\'' +
            ", order=" + price +
            '}';
    }

}
