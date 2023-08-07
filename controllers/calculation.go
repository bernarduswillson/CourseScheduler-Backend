package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singleservice/initializers"
	model "singleservice/models"
	"fmt"
)

func ScheduleCourses(c *gin.Context) {
	var requestBody struct {
		Jurusan  string `json:"jurusan"`
		Semester int    `json:"semester_ambil"`
		MinSKS   int    `json:"sks_minimal"`
		MaxSKS   int    `json:"sks_maksimal"`
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

	var courses []model.Matakuliah
	result := initializers.DB.Where("jurusan_mk = ? AND semester_minimal <= ?",
		requestBody.Jurusan, requestBody.Semester).Find(&courses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch courses",
			"data":    nil,
		})
		return
	}

	selectedCourses := courseSchedulingAlgorithm(courses, requestBody.Semester, requestBody.MinSKS, requestBody.MaxSKS)

	// total selected SKS
	totalSelectedSKS := 0
	for _, course := range selectedCourses {
		totalSelectedSKS += course.SKS
	}

	// total selected score
	totalSelectedScore := totalScore(selectedCourses)

	c.JSON(http.StatusOK, gin.H{
		"status":             "success",
		"message":            "Course schedule computed successfully",
		"selectedCourses":    selectedCourses,
		"totalSelectedSKS":   totalSelectedSKS,
		"totalSelectedScore": totalSelectedScore,
	})
}

func courseSchedulingAlgorithm(courses []model.Matakuliah, semester, minSKS, maxSKS int) []model.Matakuliah {
	selectedCourses := make([]model.Matakuliah, 0)
	// if maksSKS < minSKS, no course
	if maxSKS < minSKS {
		return selectedCourses
	}

	n := len(courses)

	// create a 2D slice to store the DP table
	// dp[i][j] represents the maximum score achievable with j SKS in the first i courses
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, maxSKS+1)
	}

	// fill the DP table using bottom-up dynamic programming
	for i := 1; i <= n; i++ {
		for j := 1; j <= maxSKS; j++ {
			if courses[i-1].SemesterMinimal <= semester {
				if courses[i-1].SKS <= j {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-courses[i-1].SKS]+mapPrediksiToScore(courses[i-1].Prediksi)*float64(courses[i-1].SKS))
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// print the DP table
	fmt.Println("DP Table:")
	for i := 0; i <= n; i++ {
		for j := 0; j <= maxSKS; j++ {
			fmt.Printf("%.2f ", dp[i][j])
		}
		fmt.Println()
	}

	// backtrack to find the selected courses based on the DP table
	i := n
	j := maxSKS
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			selectedCourses = append(selectedCourses, courses[i-1])
			j -= courses[i-1].SKS
		}
		i--
	}

	// reverse the selectedCourses
	reverse(selectedCourses)

	return selectedCourses
}

func totalScore(courses []model.Matakuliah) float64 {
	totalSKS := 0
	total := 0.0

	for _, course := range courses {
		totalSKS += course.SKS
		prediksiScore := mapPrediksiToScore(course.Prediksi)
		total += prediksiScore * float64(course.SKS)
	}

	if totalSKS == 0 {
		return 0.0
	}

	return total / float64(totalSKS)
}

func mapPrediksiToScore(prediksi string) float64 {
	switch prediksi {
	case "A":
		return 4.0
	case "AB":
		return 3.5
	case "B":
		return 3.0
	case "BC":
		return 2.5
	case "C":
		return 2.0
	case "D":
		return 1.0
	default:
		return 0.0
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func reverse(s []model.Matakuliah) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
