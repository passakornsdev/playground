package utils

import (
	"github.com/go-resty/resty/v2"
)

var WEATHER_BASE_URL string = "http://localhost:8090"
var client = resty.New()

type WeatherResponse struct {
	Temp int8 `json="temp"`
}

func GetTemperature(cityName string) (int8, error) {

	var responsebody WeatherResponse
	_, err := client.R().
		SetQueryParams(map[string]string{
			"cityName": cityName,
		}).
		SetResult(&responsebody).
		Get(WEATHER_BASE_URL + "/v1/api/weather")

	if err != nil {
		return 0, err
	}

	return responsebody.Temp, nil
}
