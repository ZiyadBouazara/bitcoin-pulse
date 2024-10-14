package com.pricefetcherservice.infrastructure.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.util.List;


@Getter
@Setter
@ToString
public class SubscribePriceDTO {
    @JsonProperty("type")
    private String type;

    @JsonProperty("channels")
    private List<String> channels;

    @JsonProperty("product_ids")
    private List<String> stocks;

    public SubscribePriceDTO(List<String> stocks) {
        this.type = "subscribe";
        this.channels = List.of("ticker");
        this.stocks = stocks;
    }
}