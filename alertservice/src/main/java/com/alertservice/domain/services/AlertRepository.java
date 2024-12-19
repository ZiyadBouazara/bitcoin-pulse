package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;

import java.math.BigDecimal;
import java.util.List;

public interface AlertRepository {
    void saveAlert(Alert alert);
    Alert getAlert(Long alertId);
    void updateAlert(Long id, BtcPrice newPrice, BigDecimal newPriceDifference);
    void deleteAlert(Long alertId);
    List<Alert> getAlerts(String triggerPrice);
}
