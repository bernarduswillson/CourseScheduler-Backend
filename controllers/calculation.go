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
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, maxSKS+1)
    }

    // Fill the DP table using bottom-up dynamic programming
    for i := 1; i <= n; i++ {
        for j := minSKS; j <= maxSKS; j++ {
            // Convert SKS to int before performing operations
            sks := int(courses[i-1].SKS)

            if sks > j {
              dp[i][j] = dp[i-1][j]
            } else {
                // Convert Prediksi to int score using the mapping function
                prediksiScore := mapPrediksiToScore(courses[i-1].Prediksi)
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-sks]+prediksiScore)
            }
        }
    }

    // Backtrack to find the selected courses based on the DP table
    selectedCourses := make([]model.Matakuliah, 0)
    i, j := n, maxSKS
    for i > 0 && j >= minSKS {
        // Convert SKS to int before performing operations
        sks := int(courses[i-1].SKS)

        // If the score at dp[i][j] is different from dp[i-1][j], it means the course is selected
        if dp[i][j] != dp[i-1][j] {
            selectedCourses = append(selectedCourses, courses[i-1])
            j -= sks
        }
        i--
    }

    // Reverse the selectedCourses since we added them in reverse order during backtracking
    reverse(selectedCourses)

    return selectedCourses
}

func mapPrediksiToScore(prediksi string) int {
  switch prediksi {
  case "A":
      return 4
  case "AB":
      return 3
  case "B":
      return 3
  case "BC":
      return 2
  case "C":
      return 2
  case "D":
      return 1
  default:
      return 0
  }
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
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

