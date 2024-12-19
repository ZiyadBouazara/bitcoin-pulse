package com.alertservice.infrastructure;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;
import com.alertservice.domain.services.AlertRepository;

import java.math.BigDecimal;
import java.util.List;

public class InDBAlertRepository implements AlertRepository {
    @Override
    public void saveAlert(Alert alert) {
    }

    @Override
    public Alert getAlert(Long alertId) {
        return null;
    }

    @Override
    public void updateAlert(Long id, BtcPrice newPrice, BigDecimal newPriceDifference) {
        // todo: add inmemo + db connection here
    }

    @Override
    public void deleteAlert(Long alertId) {
    }

    @Override
    public List<Alert> getAlerts(String triggerPrice) {
        return null;
    }
}
