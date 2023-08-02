package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singleservice/initializers"
	model "singleservice/models"
	"fmt"
)

func AddMatkul (c *gin.Context) {
	fmt.Println("AddMatkul")
	var requestBody struct {
		NamaMK string `json:"nama_mk"`
		SKS    uint    `json:"sks"`
		JurusanMK string `json:"jurusan_mk"`
		SemesterMinimal uint `json:"semester_minimal"`
		Prediksi string `json:"prediksi"`
	}

	err := c.BindJSON(&requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	newMatkul := model.Matakuliah{
		NamaMK: requestBody.NamaMK,
		SKS: requestBody.SKS,
		JurusanMK: requestBody.JurusanMK,
		SemesterMinimal: requestBody.SemesterMinimal,
		Prediksi: requestBody.Prediksi,
	}

	result := initializers.DB.Create(&newMatkul)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create matakuliah",
			"data":    nil,
		})
		return
	}

	data := gin.H{
		"id": newMatkul.ID,
		"nama_mk": newMatkul.NamaMK,
		"sks": newMatkul.SKS,
		"jurusan_mk": newMatkul.JurusanMK,
		"semester_minimal": newMatkul.SemesterMinimal,
		"prediksi": newMatkul.Prediksi,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Matakuliah created successfully",
		"data":    data,
	})
}