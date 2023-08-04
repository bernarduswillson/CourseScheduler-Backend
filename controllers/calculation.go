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
  // Helper function to find all possible course combinations that meet the SKS constraints
  var findCourseCombinations func(int, float64, float64, []model.Matakuliah, []model.Matakuliah)

  selectedCourses := make([]model.Matakuliah, 0)
  currentCourses := make([]model.Matakuliah, 0)

  findCourseCombinations = func(index int, currentSKS, currentScore float64, coursesList, currentList []model.Matakuliah) {
      if currentSKS >= float64(minSKS) && currentSKS <= float64(maxSKS) {
          if currentScore > totalScore(selectedCourses) {
              selectedCourses = append([]model.Matakuliah{}, currentList...)
          }
      }

      if index >= len(courses) {
          return
      }

      findCourseCombinations(index+1, currentSKS, currentScore, coursesList, currentList)

      course := courses[index]
      sks := float64(course.SKS)
      prediksiScore := mapPrediksiToScore(course.Prediksi)

      currentSKS += sks
      currentScore += prediksiScore
      currentList = append(currentList, course)

      findCourseCombinations(index+1, currentSKS, currentScore, coursesList, currentList)
  }

  findCourseCombinations(0, 0, 0, courses, currentCourses)
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
