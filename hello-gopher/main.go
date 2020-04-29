package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gopher",
		})
	})
	
	fmt.Println("Testing")

	// for _, var := range var {
	// 	fmt.Println("Testing")
	// }

	r.Run()
}
