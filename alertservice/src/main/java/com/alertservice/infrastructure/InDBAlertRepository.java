package com.alertservice.infrastructure;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.services.AlertRepository;

public class InDBAlertRepository implements AlertRepository {
    @Override
    public void saveAlert(Alert alert) {

    }

    @Override
    public Alert getAlert(String alertId) {
        return null;
    }

    @Override
    public void updateAlert(Alert alert) {

    }

    @Override
    public void deleteAlert(String alertId) {

    }
}
