package student

import (
	"awesomeProject1/dal/db"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryStudentInfo(c *gin.Context) {
	var login *GetStudentReq
	var stu *model.Student
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": err.Error()})
		return
	}
	if login.Opt == "SelectById" {
		stu, err := db.SelectByID(login.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
		}
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
}
