package com.alertservice.domain.services;

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

}
