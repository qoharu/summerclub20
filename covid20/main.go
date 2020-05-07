package main

import (
	"summerclub20/covid20/controller"
	"summerclub20/covid20/database"
	"summerclub20/covid20/migration"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initialize database and application.
	database.InitDB()
	migration.Migrate()
	r := gin.Default()

	// Router
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gopher",
		})
	})

	r.GET("/daily/getByDate", controller.GetDailyUpdateByDate)

	// Finally, run the application
	r.Run()
}
