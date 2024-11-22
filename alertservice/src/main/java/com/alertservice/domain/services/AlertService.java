package com.alertservice.domain.services;

import org.springframework.stereotype.Service;

@Service
public class AlertService {
    private PriceFilter priceFilter;

    public AlertService(PriceFilter priceFilter) {
        this.priceFilter = priceFilter;
    }

}
