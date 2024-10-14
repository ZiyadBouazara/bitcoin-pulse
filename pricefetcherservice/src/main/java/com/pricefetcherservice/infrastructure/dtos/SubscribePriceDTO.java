package com.pricefetcherservice.infrastructure.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.Getter;
import lombok.Setter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;


@Getter
@Setter
public class SubscribePriceDTO {
    private static final Logger logger = LoggerFactory.getLogger(SubscribePriceDTO.class);
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

    @Override
    public String toString() {
        try {
            return new ObjectMapper().writeValueAsString(this);
        } catch (JsonProcessingException e) {
            logger.error("Error converting SubscribePriceDTO to JSON string", e);
            throw new RuntimeException(e);
        }
    }
}