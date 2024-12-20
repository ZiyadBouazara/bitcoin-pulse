package com.alertservice.domain.services;

import com.alertservice.domain.models.Alert;
import software.amazon.awssdk.services.sns.SnsClient;
import software.amazon.awssdk.services.sns.model.PublishRequest;
import software.amazon.awssdk.services.sns.model.PublishResponse;

import java.util.List;

public class SNSAlertSender implements AlertSender {

    private SnsClient snsClient;

    public SNSAlertSender() {
        this.snsClient = SnsClient.create();
    }

    @Override
    public void sendAlerts(List<Alert> alerts) {
        for (Alert alert : alerts) {
            try {
                String message = formatAlertMessage(alert);
                PublishRequest publishRequest = PublishRequest.builder()
                        .message(message)
                        .phoneNumber(alert.phone()) // Optional, if sending SMS
                        .subject("BTC Price Alert")
                        .build();

                PublishResponse publishResponse = snsClient.publish(publishRequest);
                System.out.println("Alert sent for email: " + alert.email() +
                        ", MessageId: " + publishResponse.messageId());
            } catch (Exception e) {
                System.err.println("Failed to send alert to email: " + alert.email() +
                        ", Error: " + e.getMessage());
            }
        }
    }

    private String formatAlertMessage(Alert alert) {
        return "BTC reached you price target!\n"
                + "Target Price: " + alert.triggerPrice() + "\n"
                + "Your Details: " + alert.message();
    }
}
