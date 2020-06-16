package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()

	// Akses di: localhost:8080/hello?fullName=Raven&helloMessage=Hai&birthYear=1994
	router.GET("/hello", func(context *gin.Context) {
		var fullName string = context.Query("fullName")
		var helloMessage string = context.DefaultQuery("helloMessage", "Good Morning")

		var birthYear, _ = strconv.Atoi(context.Query("birthYear"))
		var age = 2020 - birthYear

		context.JSON(200, gin.H{
			"result": gin.H{
				"fullName":        fullName,
				"helloMessage":    helloMessage,
				"combinedMessage": helloMessage + " " + fullName + " Your age is " + strconv.Itoa(age), // strconv.Itoa will convert from integer to alphabet/string
			},
			"success":      true,
			"errorMessage": nil,
		})
	})

	// Akses di: localhost:8080/whoami?fullName=Raven
	router.GET("/whoami", func(context *gin.Context) {
		var fullName string = context.DefaultQuery("fullName", "Guest")

		context.JSON(200, gin.H{
			"result": gin.H{
				"message":  "You're " + fullName,
				"fullName": fullName,
			},
			"success":      true,
			"errorMessage": nil,
		})
	})

	router.Run()
}
