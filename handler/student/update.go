package student

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateStudentInfo(c *gin.Context) {
	var req *UpdateStudentInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "error", err.Error())
		return
	}
	stu, err := db.UpdateStudentInformation(&model.Student{
		ID:          uint(req.ID),
		Sex:         req.Sex,
		Name:        req.Name,
		Grade:       req.Grade,
		StudentType: req.StudentType,
	})
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, stu)
	}
	reply.Reply(c, http.StatusOK, "Update success", stu)
}
