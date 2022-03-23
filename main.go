package main

import (
	Db "awesomeProject1/dal"
	"awesomeProject1/handler/Graduate"
	"awesomeProject1/handler/Selection"
	"awesomeProject1/handler/course"
	"awesomeProject1/handler/student"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	Db.Init()
}
func main() {
	Init()
	r := gin.Default()
	r.GET("/student", student.QueryStudentInfo)
	r.POST("/student", student.UpdateStudentInfo)
	r.PUT("/student", student.CreateStudentInfo)
	r.DELETE("/student", student.DeleteStudentInfo)

	r.GET("/course", course.QueryCourseInfo)
	r.POST("/course", course.UpdateCourseInfo)
	r.PUT("/course", course.CreateCourseInfo)
	r.DELETE("/course", course.DeleteCourseInfo)

	r.PUT("/selection", Selection.InsertCourseSelectionInfo)
	r.GET("/selection", Selection.GetCourseSelectionInfo)
	r.DELETE("/selection", Selection.CancelCourseSelection)

	r.GET("/graduate", Graduate.Graduate)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
