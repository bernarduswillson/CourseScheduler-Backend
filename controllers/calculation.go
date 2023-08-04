package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singleservice/initializers"
	model "singleservice/models"
)

func ScheduleCourses(c *gin.Context) {
	var requestBody struct {
		Jurusan     string `json:"jurusan"`
		Semester    int    `json:"semester_ambil"`
		MinSKS      int    `json:"sks_minimal"`
		MaxSKS      int    `json:"sks_maksimal"`
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
	result := initializers.DB.Where("jurusan_mk = ? AND semester_minimal <= ? AND sks >= ? AND sks <= ?",
		requestBody.Jurusan, requestBody.Semester, requestBody.MinSKS, requestBody.MaxSKS).Find(&courses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch courses",
			"data":    nil,
		})
		return
	}

	selectedCourses := courseSchedulingAlgorithm(courses, requestBody.Semester, requestBody.MinSKS, requestBody.MaxSKS)

	c.JSON(http.StatusOK, gin.H{
		"status":          "success",
		"message":         "Course schedule computed successfully",
		"selectedCourses": selectedCourses,
	})
}

func courseSchedulingAlgorithm(courses []model.Matakuliah, semester, minSKS, maxSKS int) []model.Matakuliah {
	n := len(courses)

	// Create a 2D slice to store the DP table
	// dp[i][j] represents the maximum score achievable with j SKS in the first i courses
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, maxSKS+1)
	}

	// Fill the DP table using bottom-up dynamic programming
	for i := 1; i <= n; i++ {
		for j := minSKS; j <= maxSKS; j++ {
			// Convert SKS to float64 before performing operations
			sks := float64(courses[i-1].SKS)

			if sks > float64(j) {
				dp[i][j] = dp[i-1][j]
			} else {
				// Convert Prediksi to float64 score using the mapping function
				prediksiScore := mapPrediksiToScore(courses[i-1].Prediksi)
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-int(sks)]+prediksiScore)
			}
		}
	}

	// Backtrack to find the selected courses based on the DP table
	selectedCourses := make([]model.Matakuliah, 0)
	i, j := n, maxSKS
	for i > 0 && j >= minSKS {
		// Convert SKS to float64 before performing operations
		sks := float64(courses[i-1].SKS)

		// If the score at dp[i][j] is different from dp[i-1][j], it means the course is selected
		if dp[i][j] != dp[i-1][j] {
			selectedCourses = append(selectedCourses, courses[i-1])
			j -= int(sks)
		}
		i--
	}

	// Reverse the selectedCourses since we added them in reverse order during backtracking
	reverse(selectedCourses)

	return selectedCourses
}

// Helper function to calculate the total score of a list of courses
func totalScore(courses []model.Matakuliah) float64 {
	total := 0.0
	for _, course := range courses {
		prediksiScore := mapPrediksiToScore(course.Prediksi)
		total += prediksiScore
	}
	return total
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

// Helper function to find the maximum of two floats
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Helper function to reverse a slice of Matakuliah
func reverse(s []model.Matakuliah) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
