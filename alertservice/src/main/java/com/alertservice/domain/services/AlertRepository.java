package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;

public interface AlertRepository {
    void saveAlert(Alert alert);
    Alert getAlert(String alertId);
    void updateAlert(Alert alert);
    void deleteAlert(String alertId);
}
