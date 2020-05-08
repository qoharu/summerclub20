package controller

import (
	"summerclub20/covid20/database"
	"summerclub20/covid20/model"
	"time"

	"github.com/gin-gonic/gin"
)

//GetDailyUpdateByDate ... get daily update data by date
func GetDailyUpdateByDate(context *gin.Context) {
	//caution : oddly, the format string is `2006-01-02 15:04:05.000000000`
	currentDate := time.Now().Format("2006-01-02")

	date := context.DefaultQuery("date", currentDate)

	var result model.Daily
	record := database.DBCon.Where("date = ?", date).First(&result)

	status := "OK"
	if record.RecordNotFound() {
		status = "Not Found"
	}

	context.JSON(200, gin.H{
		"data":   result,
		"status": status,
		"error":  record.Error,
	})
}

//InsertDailyUpdate ... insert new records for daily update
func InsertDailyUpdate(context *gin.Context) {
	var spec model.Daily
	// context.BindJSON(&spec)

	if err := context.ShouldBindJSON(&spec); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insert := database.DBCon.Create(&spec)

	status := "OK"
	if insert.RowsAffected == 0 {
		status = "Failed"
	}

	context.JSON(200, gin.H{
		"status": status,
		"error":  insert.Error,
	})
}
