package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourseInfo(c *gin.Context) {
	var req CreateCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cou, err := db.InsertCourseInformation(&model.Course{
		CourseID:          req.CourseID,
		CourseName:        req.CourseName,
		CourseCredit:      req.CourseCredit,
		CourseNumberLimit: req.CourseNumberLimit,
		CourseType:        req.CourseType,
	})
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, cou)
	}
	reply.Reply(c, http.StatusOK, "Create success", cou)
}
