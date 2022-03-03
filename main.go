package main

import (
	Db "awesomeProject1/dal"
	"awesomeProject1/handler/courseSelection"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	Db.Init()
}
func main() {
	Init()
	r := gin.Default() /*
			r.GET("/:id/:name/:sex/:grade", student.QueryStudentInfo)
			r.POST("/student", student.UpdateStudentInfo)
			r.PUT("/student", student.CreateStudentInfo)
			r.DELETE("/student", student.DeleteStudentInfo)


		r.GET("/:courseID/:courseName/:courseCredit", course.QueryCourseInfo)
		r.POST("/course", course.UpdateCourseInfo)
		r.PUT("/course", course.CreateCourseInfo)
		r.DELETE("/course", course.DeleteCourseInfo)
	*/
	r.PUT("selection", courseSelection.InsertCourseSelectionInfo)
	r.GET("/:studentID/:classID", courseSelection.GetCourseSelectionInfo)
	r.DELETE("selection", courseSelection.CancelCourseSelection)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
