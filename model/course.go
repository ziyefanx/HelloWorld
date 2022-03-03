package model

type Course struct {
	CourseID     int    `form:"courseID" json:"courseID" uri:"courseID" xml:"courseID" binding:"required"`
	CourseName   string `form:"courseName" json:"courseName" uri:"courseName" xml:"courseName" binding:"required"`
	CourseCredit int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" xml:"courseCredit" binding:"required"`
}
