package graduate

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"awesomeProject1/model"
	"awesomeProject1/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Graduate(c *gin.Context) {
	var req GetGraduateReq
	if err := c.ShouldBindQuery(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "error", err.Error())
		return
	}
	student, err := db.SelectStudentByID(uint(req.StudentID))
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
		return
	}

	courses, err := db.GetCourseSelectionInformation(req.StudentID)
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
		return
	}

	var majorCredit, PECredit, artCredit int
	for _, course := range courses {
		switch course.CourseType {
		case model.CourseTypeArt:

		}
	}

	studentHandler, err := service.GetStudentHandler(student.StudentType)
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
		return
	}
	satisfyGraduate := studentHandler.SatisfyGraduate(majorCredit, artCredit, PECredit)

	c.JSON(http.StatusOK, gin.H{"专业课学分": majorCredit, "艺术课学分": artCredit, "体育课学分": PECredit, "能否毕业": satisfyGraduate})
}
