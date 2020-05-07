package controller

import (
	"qoharu/summerclub20/covid19/database"
	"qoharu/summerclub20/covid19/model"

	"github.com/gin-gonic/gin"
)

//GetDailyUpdateByDate ... get daily update data by date
func GetDailyUpdateByDate(c *gin.Context) {
	date := c.Query("date")

	var result model.Daily
	database.DBCon.Where("date = ?", date).First(&result)

	c.JSON(200, gin.H{
		"data":   result,
		"status": "OK",
	})
}
