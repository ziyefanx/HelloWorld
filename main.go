package main

import (
	Db "awesomeProject1/dal"
	"awesomeProject1/handler/student"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()

	r := gin.Default()
	r.GET("/student", student.QueryStudentInfo)
	r.POST("/student", student.UpdateStudentInfo)
	r.PUT("/student", student.CreateStudentInfo)
	r.DELETE("/student", student.DeleteStudentInfo)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Init() {
	Db.Init()
}
