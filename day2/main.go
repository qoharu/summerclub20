package main

import (
	"github.com/gin-gonic/gin"
)

type arithModel struct {
	A	int `json:"a" binding:"required"`
	B	int `json:"b" binding:"required"`
}

func main() {
	router := gin.Default()
	
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gopher",
		})
	})

	router.GET("/addition", func(c *gin.Context) {
		var spec arithModel
		
		err := c.ShouldBindJSON(&spec)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(200, gin.H{
			"data": spec.A + spec.B,
		})
	})

	router.POST("/substraction", func(c *gin.Context) {
		var spec arithModel
		
		err := c.ShouldBindJSON(&spec)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"data": spec.A - spec.B,
		})
	})

	router.Run()
}
