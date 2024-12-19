package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;

@Service
public class AlertService {
    private final PriceFilter priceFilter;
    private final AlertRepository alertRepository;
    private final AlertFactory alertFactory;
    private final AlertSender alertSender;

    public AlertService(PriceFilter priceFilter, AlertRepository alertRepository, AlertFactory alertFactory,
                        AlertSender alertSender) {
        this.priceFilter = priceFilter;
        this.alertFactory = alertFactory;
        this.alertRepository = alertRepository;
        this.alertSender = alertSender;
    }

    public void createAlert(BtcPrice price, BigDecimal priceDifference, String email, String phoneNumber) {
        Alert alert = alertFactory.createPriceAlert(price, priceDifference);
        alertRepository.saveAlert(alert);
    }

    public void updateAlert(Long id, BtcPrice newPrice, BigDecimal newPriceDifference) {
        alertRepository.updateAlert(id, newPrice, newPriceDifference);
    }

    public void deleteAlert(Long id) {
        alertRepository.deleteAlert(id);
    }

    public void sendTriggeredAlerts(String price) {
        List<Alert> alerts = alertRepository.getAlerts(price);
        alertSender.sendAlerts(alerts);
    }
}
