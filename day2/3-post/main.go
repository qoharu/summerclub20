package main

import (
	"github.com/gin-gonic/gin"
)

type studentModel struct {
	FullName string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Grade    int    `json:"grade" binding:"required"`
}

func main() {
	var router = gin.Default()

	router.POST("/student/add", func(context *gin.Context) {
		// inisialisasi variable
		var request studentModel

		// Masukan variable `request` ke funcion ShouldBindJSON nanti akan terisi dengan parameter pada request.
		// Fungsi ini juga akan mengembalikan error jika ada
		var bindError = context.ShouldBindJSON(&request)

		// Cek apakah variable bindError terisi, jika ya kirim response error
		if bindError != nil {
			context.JSON(400, gin.H{
				"result":       nil,
				"success":      false,
				"errorMessage": bindError.Error(),
			})
		} else {
			// Jika tidak ada error, kirim response berhasil
			context.JSON(201, gin.H{
				"result": gin.H{
					"message": "Student with name " + request.FullName + " is successfully registered",
				},
				"success":      true,
				"errorMessage": nil,
			})
		}
	})

	router.Run()
}
