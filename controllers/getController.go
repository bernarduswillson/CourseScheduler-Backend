package controllers
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singleservice/initializers"
	model "singleservice/models"
	// "fmt"
)

func GetMatkul(c *gin.Context) {
	var matkul []model.Matakuliah
	result := initializers.DB.Find(&matkul)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch matakuliah",
			"data":    nil,
		})
		return
	}

	data := []gin.H{}

	for _, v := range matkul {
		data = append(data, gin.H{
			"id": v.ID,
			"nama_mk": v.NamaMK,
			"sks": v.SKS,
			"jurusan_mk": v.JurusanMK,
			"semester_minimal": v.SemesterMinimal,
			"prediksi": v.Prediksi,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully fetch matakuliah",
		"data":    data,
	})
}