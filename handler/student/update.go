package student

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateStudentInfo(c *gin.Context) {
	var req *model.Student
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stu, err := db.UpdateStudentInformation(&model.Student{
		ID:    uint(req.ID),
		Sex:   req.Sex,
		Name:  req.Name,
		Grade: req.Grade,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": stu})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Update success", "data": stu})
}
