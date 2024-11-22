package com.alertservice.domain.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.annotation.Nullable;

public record Alert(
        @JsonProperty("type")
        String type,
        @JsonProperty("message")
        String message,
        @JsonProperty("email")
        @Nullable
        String email,
        @JsonProperty("phone")
        @Nullable
        String phone,
        @JsonProperty("price")
        BtcPrice bitcoinPrice,
        @JsonProperty("direction")
        TrendDirection direction
) {


}
