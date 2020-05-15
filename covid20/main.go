package main

import (
	"covid20/controller"
	"covid20/database"
	"covid20/migration"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initialize database and application.
	database.InitDB()
	migration.Migrate()
	router := gin.Default()

	router.GET("/daily/getByDate", controller.GetDailyUpdateByDate)
	
	router.POST("/daily/insert", controller.InsertDailyUpdate)

	// Finally, run the application
	router.Run()
}