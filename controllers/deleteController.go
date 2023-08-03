package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singleservice/initializers"
	model "singleservice/models"
	"strconv"
)

func DeleteMatkul(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid ID format",
			"data":    nil,
		})
		return
	}

	var matkul model.Matakuliah
	result := initializers.DB.First(&matkul, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to find matakuliah",
			"data":    nil,
		})
		return
	}

	result = initializers.DB.Delete(&matkul)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete matakuliah",
			"data":    nil,
		})
		return
	}

	data := gin.H{
		"id":               matkul.ID,
		"nama_mk":          matkul.NamaMK,
		"sks":              matkul.SKS,
		"jurusan_mk":       matkul.JurusanMK,
		"semester_minimal": matkul.SemesterMinimal,
		"prediksi":         matkul.Prediksi,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Matakuliah deleted successfully",
		"data":    data,
	})
}
