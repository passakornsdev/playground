package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api-test/utils"
	"github.com/gin-gonic/gin"
)

func getWeather(context *gin.Context) {
	cityName := context.Query("cityName")

	if cityName == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "not city name provided",
		})
		return
	}

	temp, err := utils.GetTemperature(cityName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error occur while trying to fetch temperature",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v now is %v degree celcius", cityName, temp),
	})
}
