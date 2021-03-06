package main

import "github.com/gin-gonic/gin"

func main() {
	var router = gin.Default()

	// Akses di localhost:8080/hello
	router.GET("/hello", func(context *gin.Context) {
		var fullName string = "Salman"
		var helloMessage string = "Good morning"

		context.JSON(200, gin.H{
			"result": gin.H{
				"fullName":     fullName,
				"helloMessage": helloMessage,
			},
			"success":      true,
			"errorMessage": nil, // nil is null, null is "no value"
		})
	})

	router.Run()
}
