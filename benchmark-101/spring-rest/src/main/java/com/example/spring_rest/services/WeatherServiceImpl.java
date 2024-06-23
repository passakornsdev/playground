package com.example.spring_rest.services;

import com.example.spring_rest.dto.weatherService.GetTemperatureResponse;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.HashMap;
import java.util.Map;

@Service
@RequiredArgsConstructor
public class WeatherServiceImpl implements WeatherService {

    private static final Logger log = LoggerFactory.getLogger(WeatherServiceImpl.class);
    @Value("${app.weather-service.base-url}")
    private String basedUrl = "";

    private final RestTemplate restClient;

    @Override
    public int getTemperature(String city) {
        log.info("Sending request to weather service");
        Map<String, String> params = new HashMap<>();
        params.put("cityName", city);
        GetTemperatureResponse response = restClient.getForObject(this.basedUrl.concat("/v1/api/weather"), GetTemperatureResponse.class, params);
        log.info("Successfully Sent request to weather service");
        return response == null ? 0 : response.getTemp();
    }
}
