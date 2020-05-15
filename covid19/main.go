package main

import (
	"covid19/controller"
	"covid19/database"
	"covid19/migration"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initialize database and application.
	database.InitDB()
	migration.Migrate()
	router := gin.Default()

	router.GET("/daily/getByDate", controller.GetDailyUpdateByDate)

	// Finally, run the application
	router.Run()
}
