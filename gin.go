package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Api    string `form:"api" json:"api" uri:"api" xml:"api" binding:"required"`
	People string `form:"people" json:"people" uri:"people" xml:"people" binding:"required"`
	Select string `form:"select" json:"select" uri:"select" xml:"select" binding:"required"`
	Name   string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Sex    uint   `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
	Grade  int    `form:"grade" json:"grade" uri:"grade" xml:"grade" binding:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.GET("/:api/:people/:select/:name/:sex/:grade", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.Api != "api" || login.People != "student" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		if login.Select == "insert" {
			InsertStudentInformation(login.Name, login.Sex, login.Grade)
			c.JSON(http.StatusOK, gin.H{"status": "200"})
		}
	})
	r.Run(":8000")
}
