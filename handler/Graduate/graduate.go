package Graduate

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Graduate(c *gin.Context) {
	var flag bool = false
	var majorCredit int = 0
	var PECredit int = 0
	var artCredit int = 0
	var req GetGraduateReq
	if err := c.ShouldBindQuery(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "error", err.Error())
		return
	}
	student, Err := db.SelectStudentByID(uint(req.StudentID))
	if Err != nil {
		reply.Reply(c, http.StatusInternalServerError, "500", Err.Error())
		return
	}
	courses := make([]model.Course, 0)
	courses, Err = db.GetCourseSelectionInformation(req.StudentID)
	for _, course := range courses {
		if course.CourseType == "专业课程" {
			majorCredit += course.CourseCredit
		} else if course.CourseType == "艺术课程" {
			artCredit += course.CourseCredit
		} else {
			PECredit += course.CourseCredit
		}
	}
	if student.StudentType == "普通学生" {
		if majorCredit >= 10 && artCredit >= 3 && PECredit >= 5 {
			flag = true
		}
	} else if student.StudentType == "艺术生" {
		if majorCredit >= 5 && artCredit >= 10 && PECredit >= 3 {
			flag = true
		}
	} else {
		if majorCredit >= 5 && artCredit >= 3 && PECredit >= 10 {
			flag = true
		}
	}
	c.JSON(http.StatusOK, gin.H{"专业课学分": majorCredit, "艺术课学分": artCredit, "体育课学分": PECredit, "能否毕业": flag})
}
