package student

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudentInfo(c *gin.Context) {
	var req *UpdateStudentInfoReq
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
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, stu))
	}
	c.JSON(http.StatusOK, db.Reply("200", "Create success", stu))
}
