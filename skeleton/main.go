package main

import (
	"summerclub20/skeleton/controller"
	"summerclub20/skeleton/database"
	"summerclub20/skeleton/migration"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initialize database and application.
	database.InitDB()
	migration.Migrate()
	router := gin.Default()

	router.GET("/daily/getByDate", controller.GetDailyUpdateByDate)

	// router.POST("/daily/insert", controller.InsertDailyUpdate)

	// Finally, run the application
	router.Run()
}
