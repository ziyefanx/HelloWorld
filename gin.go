package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Opt   string `form:"opt" json:"opt" uri:"opt" xml:"opt" binding:"required"`
	Id    int    `form:"id" json:"id" uri:"id" xml:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" xml:"grade" binding:"required"`
}

func main() {
	Init()
	Db.AutoMigrate(&Student{})
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.Opt == "Add" {
			InsertStudentInformation(json.Name, json.Sex, json.Grade)
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Add success"})
			return
		}
		if json.Opt == "Update" {
			UpdateStudentInformation(json.Name, json.Id, json.Grade, json.Sex)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Update success"})
		}
		if json.Opt == "Delete" {
			DeleteStudentInformation(json.Id)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Delete success"})
		}
	})
	r.GET("/:opt/:id/:name/:sex/:grade", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.Opt == "SelectById" {
			SelectByID(login.Id)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": stu})
		}
		if login.Opt == "SelectByName" {
			SelectByName(login.Name)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": stu})
		}
		if login.Opt == "SelectBySex" {
			SelectBySex(login.Sex)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": stu})
		}
		if login.Opt == "SelectByGrade" {
			SelectByGrade(login.Grade)
			c.JSON(http.StatusOK, gin.H{"status": "200", "massage": stu})
		}
	})
	r.Run(":8000")
}
