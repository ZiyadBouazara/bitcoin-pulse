/*
package com.alertservice;

import com.alertservice.domain.models.Alert;
import com.alertservice.domain.models.BtcPrice;
import com.alertservice.domain.services.AlertSender;
import com.alertservice.domain.services.AlertService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.springframework.boot.test.context.SpringBootTest;

import java.math.BigDecimal;

import static jdk.internal.org.objectweb.asm.util.CheckClassAdapter.verify;
import static jdk.jfr.internal.jfc.model.Constraint.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.when;

@SpringBootTest
class AlertserviceApplicationTests {

    public static final BtcPrice PRICE = new BtcPrice(BigDecimal.valueOf(10000));
    public static final BigDecimal PRICE_DIFFERENCE = BigDecimal.valueOf(100);
    public static final String EMAIL = "email@email.com";
    public static final String PHONENUMBER = "1234567890";
    @Mock
    AlertSender alertSender;

    @BeforeEach
    void setUp() {
    }

    @InjectMocks
    AlertService alertService;

    @Test
    givenCorrectParams_whenCreateAlert_thenAlertIsCreated() {
        Alert expectedAlert = new Alert(, , EMAIL, PHONENUMBER);
        when(alertRepository.saveAlert(any(Alert.class))).thenReturn(expectedAlert);

        alertService.createAlert(PRICE, PRICE_DIFFERENCE, EMAIL, PHONENUMBER);
        verify(alertRepository, times(1)).saveAlert(any(Alert.class));
    }
}
*/
