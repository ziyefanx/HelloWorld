package main

import (
	"awesomeProject1/dal/db"
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
	r.DELETE("/stduent", student.DeleteStudentInfo)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Init() {
	db.Init()
}
