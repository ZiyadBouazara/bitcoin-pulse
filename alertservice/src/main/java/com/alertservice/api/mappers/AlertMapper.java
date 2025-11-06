package com.alertservice.api.mappers;

import com.alertservice.api.requests.AlertRequest;
import com.alertservice.domain.models.Alert;

public class AlertMapper {

    public static Alert fromRequest(AlertRequest alertRequest) {
        return new Alert(
                "ALERT_TYPE_ON: " + alertRequest.getSymbol(),
                alertRequest.getText(),
                alertRequest.getEmail(),
                alertRequest.getPhone(),
                alertRequest.getPrice()
        );
    }
}
