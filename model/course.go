package model

type Course struct {
	CourseID          int `gorm:"primary_key"`
	CourseName        string
	CourseCredit      int
	CourseNumberLimit int
}
