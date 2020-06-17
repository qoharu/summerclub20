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

	router.GET("/car/getDetail", func(context *gin.Context) {
		var carId string = context.Query("carId")

		context.JSON(200, gin.H{
			"result": gin.H{
				"id":   carId,
				"name": "Nissan " + carId + " SuperCar",
				"year": 2019,
			},
			"success":      true,
			"errorMessage": nil,
		})
	})

	router.POST("/car/insertDetail", func(context *gin.Context) {
		
		type carModel struct {
			Name string `json:"name"`
			Year int `json:"year"`
			Color string `json:"color"`
		}

		var newCar carModel

		context.ShouldBindJSON(&newCar)

		context.JSON(200, gin.H{
			"result": gin.H{
				"message": "Car with name " + newCar.Name + " is created",
			},
			"success":      true,
			"errorMessage": nil,
		})
	})


	router.Run()
}
