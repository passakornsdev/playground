package com.example.spring_rest.controllers;

import com.example.spring_rest.dto.GetWeatherResponse;
import com.example.spring_rest.services.WeatherService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/weather")
@RequiredArgsConstructor
public class WeatherController {

    private final WeatherService weatherService;

    @GetMapping
    public GetWeatherResponse getWeather(@RequestParam String cityName) {
        int temp = this.weatherService.getTemperature(cityName);
        return GetWeatherResponse
                .builder()
                .message(String.format("%s now is %d degree celcius", cityName, temp))
                .build();
    }
}
