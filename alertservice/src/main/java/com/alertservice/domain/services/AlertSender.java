package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import software.amazon.awssdk.services.sns.SnsAsyncClient;
import software.amazon.awssdk.services.sns.SnsClient;

import java.util.List;

public interface AlertSender {

    void sendAlerts(List<Alert> alerts);

}
