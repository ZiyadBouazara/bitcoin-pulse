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
    private static final String SUBSCRIBE_TYPE = "subscribe";
    @JsonProperty("type")
    private String type;
    @JsonProperty("channels")
    private List<String> channels;
    @JsonProperty("product_ids")
    private List<String> stocks;

    public SubscribePriceDTO(List<String> channels, List<String> stocks) {
        this.type = SUBSCRIBE_TYPE;
        this.channels = channels;
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