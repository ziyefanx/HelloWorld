package student

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudentInfo(c *gin.Context) {
	var req *model.Student
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stu, err := db.InsertStudentInformation(&model.Student{
		ID:    req.ID,
		Sex:   req.Sex,
		Name:  req.Name,
		Grade: req.Grade,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": stu})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "create success", "data": stu})
}
