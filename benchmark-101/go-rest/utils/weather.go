package utils

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

var WEATHER_BASE_URL string = os.Getenv("WEATHER_BASE_URL")
var client = resty.New()

type WeatherResponse struct {
	Temp int8 `json="temp"`
}

func GetTemperature(cityName string) (int8, error) {

	var responsebody WeatherResponse

	fmt.Printf("sending request to %s \n", WEATHER_BASE_URL+"/v1/api/weather")

	response, err := client.R().
		SetQueryParams(map[string]string{
			"cityName": cityName,
		}).
		SetResult(&responsebody).
		Get(WEATHER_BASE_URL + "/v1/api/weather")

	fmt.Printf("Receiving response with status %v\n", response.Status())
	if err != nil {
		return 0, err
	}

	return responsebody.Temp, nil
}
