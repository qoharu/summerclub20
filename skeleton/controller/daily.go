package controller

import (
	"summerclub20/skeleton/model"

	"github.com/gin-gonic/gin"
)

//GetDailyUpdateByDate ... get daily update data by date
func GetDailyUpdateByDate(context *gin.Context) {
	var result model.Daily
	context.JSON(200, gin.H{
		"data":   result,
		"status": "OK",
		"error":  nil,
	})
}
