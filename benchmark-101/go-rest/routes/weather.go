package routes

import (
	"fmt"
	"net/http"

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

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v: 20 degree", cityName),
	})
}
