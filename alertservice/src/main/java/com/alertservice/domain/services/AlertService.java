package com.alertservice.domain.services;

import ch.qos.logback.core.net.server.Client;
import com.alertservice.domain.models.Alert;
import com.pricefetcherservice.domain.models.PriceEvent;
import org.springframework.stereotype.Service;

@Service
public class AlertService {
    private PriceFilter priceFilter;
    private AlertRepository alertRepository;

    public AlertService(PriceFilter priceFilter, AlertRepository alertRepository) {
        this.priceFilter = priceFilter;
        this.alertRepository = alertRepository;
    }

    public void createAlert(Alert alert) {
        alertRepository.saveAlert(alert);
    }

    public void updateAlert(Long id, Alert alert) {
        Alert alert = alertRepository.getAlert(id);
        alertRepository.updateAlert(alert);
    }

    public void deleteAlert(Long id) {
        alertRepository.deleteAlert(id);
    }
}
