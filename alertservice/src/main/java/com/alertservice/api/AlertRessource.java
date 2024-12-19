package com.alertservice.api;

import com.alertservice.api.mappers.AlertMapper;
import com.alertservice.api.requests.AlertRequest;
import com.alertservice.domain.models.Alert;
import com.alertservice.domain.services.AlertService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/alerts")
public class AlertRessource {

    @Autowired
    private AlertService alertService;

    @PostMapping
    public void createAlert(@RequestBody AlertRequest alertRequest) {
        alertService.createAlert(
                alertRequest.getPrice(),
                alertRequest.getPriceDifference(),
                alertRequest.getEmail(),
                alertRequest.getPhone()
                );
    }

    @PutMapping("/{id}")
    public void updateAlert(@PathVariable Long id, @RequestBody AlertRequest alertRequest) {
        Alert alert = AlertMapper.fromRequest(alertRequest);
        alertService.updateAlert(id, alertRequest.getPrice(), alertRequest.getPriceDifference());
    }

    @DeleteMapping("/{id}")
    public void deleteAlert(@PathVariable Long id) {
        alertService.deleteAlert(id);
    }
}
