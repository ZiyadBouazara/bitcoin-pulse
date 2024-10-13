package com.pricefetcherservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.client.RestTemplate;
import org.springframework.context.annotation.Bean;


@SpringBootApplication
public class PriceFetcherServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(PriceFetcherServiceApplication.class, args);
    }

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
