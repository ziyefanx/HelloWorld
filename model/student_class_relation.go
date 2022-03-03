package model

import "github.com/jinzhu/gorm"

type StudentCourseRelation struct {
	gorm.Model
	StudentID int `form:"studentID" json:"studentID" uri:"studentID" xml:"studentID" binding:"required"`
	ClassID   int `form:"classID" json:"classID" uri:"classID" xml:"classID" binding:"required"`
}
