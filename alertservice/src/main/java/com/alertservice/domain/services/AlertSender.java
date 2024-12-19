package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import software.amazon.awssdk.services.sns.SnsAsyncClient;
import software.amazon.awssdk.services.sns.SnsClient;

import java.util.List;

public class AlertSender {
    SnsClient sns_client = SnsClient.create();
    SnsAsyncClient sns_async_client = SnsAsyncClient.create();

    // todo: implement AWS SNS sender or any other service here.

    void sendAlerts(List<Alert> alerts) {
    }

}
