package com.example.spring_rest.dto.weatherService;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class GetTemperatureResponse {

    private int temp;
}
