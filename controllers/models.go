// Package controllers provides functions for managing students, courses, and scores.
package controllers

// Student represents a student.
type Student struct {
	// gorm.Model
	Name   string `json:"name"`   // Student's name
	Mobile string `json:"mobile"` // Student's mobile number
	Email  string `json:"email"`  // Student's email address
}

// Course represents a course.
type Course struct {
	// gorm.Model
	CourseName string `json:"course_name"` // Course name
}

// StudentScore represents a student_score
type StudentScore struct {
	// gorm.Model
	Student int    `json:"student"` // Student ID
	Course  int    `json:"course"`  // Course ID
	Score   string `json:"score"`   // Student's score
}

// TableName returns the table name for the Student struct.
func (s *Student) TableName() string {
	return "students"
}

// TableName returns the table name for the Course struct.
func (s *Course) TableName() string {
	return "courses" // Consider using "courses" for consistency.
}

// TableName returns the table name for the StudentScore struct.
func (s *StudentScore) TableName() string {
	return "studentscores" // Consider using "studentscores" to follow plural convention.
}
