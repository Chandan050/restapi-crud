package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateStudent creates a new student
// @Summary Create a new student
// @Description Create a new student with the provided details
// @Accept json
// @Produce json
// @Param student body Student true "Student details"
// @Success 200 {object} Student "Student created successfully"
// @Failure 400 {object}  "Invalid request"
// @Failure 500 {object}  "Internal server error"
// @Router /api/students [post]
func CreateStudent(c *gin.Context) {
	var student Student

	if err := c.BindJSON(&student); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := Db.Create(&student)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": student})
}

// CreateCourse creates a new course
// @Summary Creates a new course
// @Description Creates a new course with the provided details
// @Accept json
// @Produce json
// @Param course body Courses true "Course details"
// @Success 200 {object} Courses "Course created successfully"
// @Failure 400 {object}  "Invalid request"
// @Failure 500 {object}  "Internal server error"
// @Router /api/courses [post]
func CreateCourse(c *gin.Context) {
	var course Course
	if err := c.BindJSON(&course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := Db.Create(&course)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": course})
}

// CreateScore creates a new student score and subject details
// @Summary Creates a new student score
// @Description Creates a new student score with the provided details
// @Accept json
// @Produce json
// @Param studentScore body StudentScore true "Student Score details"
// @Success 200 {object} StudentScore "Student score created successfully"
// @Failure 400 {object} "Invalid request"
// @Failure 500 {object} "Internal server error"
// @Router /api/scores [post]
func CreateScore(c *gin.Context) {
	var score StudentScore
	if err := c.BindJSON(&score); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := Db.Create(&score)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": score})
}

// Result represents a score result
type Result struct {
	CourseName string `json:"course_name"`
	Score      string `json:"score"`
}

// GetScore gets all scores for a student
// @Summary Get all scores for a student
// @Description Get all scores for a student with the provided ID
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {array} Result "Scores retrieved successfully"
// @Failure 400 {object}  "Invalid request"
// @Failure 404 {object}  "Record not found"
// @Failure 500 {object}  "Internal server error"
// @Router /api/students/{id}/scores [get]
func GetScore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var scores []Result
	record := Db.Model(&StudentScore{}).
		Select("course.course_name, studentscore.score").
		Joins("JOIN courses ON studentscore.course_id = courses.ID").
		Where("studentscore.student_id = ?", id).
		Find(&scores)

	if record.Error != nil {
		if errors.Is(record.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": scores})
}

// UpdateScore updates the scores for a student
// @Summary Update scores for a student
// @Description Update scores for a student with the provided ID and Course ID
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param scoreid path int true "Course ID"
// @Param score body StudentScore true "Updated score details"
// @Success 200 {object} StudentScore "Scores updated successfully"
// @Failure 400 {object}  "Invalid request"
// @Failure 404 {object}  "Record not found"
// @Failure 500 {object}  "Internal server error"
// @Router /api/score/{id}/{scoreid} [put]
func UpdateScore(c *gin.Context) {
	var score StudentScore
	id, _ := strconv.Atoi(c.Param("id"))
	courseID, _ := strconv.Atoi(c.Param("scoreid"))

	if err := c.BindJSON(&score); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentScore StudentScore
	record := Db.Where("student_id = ? AND course_id = ?", id, courseID).First(&studentScore)
	if record.Error != nil {
		if errors.Is(record.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}

	studentScore.Score = score.Score
	Db.Save(&studentScore)
	c.JSON(http.StatusOK, gin.H{"message": studentScore})
}
