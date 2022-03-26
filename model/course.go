package model

type CourseType string

const (
	CourseTypeArt CourseType = "art"
	CourseTypePE  CourseType = "PE"
)

type Course struct {
	CourseID          int `gorm:"primary_key"`
	CourseName        string
	CourseCredit      int
	CourseNumberLimit int
	CourseType        string
}
