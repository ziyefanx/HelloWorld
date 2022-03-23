package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"awesomeProject1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// happy path 不缩进

func QueryCourseInfo(c *gin.Context) {
	var req GetCourseReq
	//var cou *model.Course
	courses := make([]model.Course, 0)
	var err error
	if err = c.ShouldBindQuery(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "500", err.Error())
		return
	}
	if req.CourseName != "" {
		courses, err = db.SelectCourseByName(req.CourseName, req.Page, req.Size)
		if err != nil {
			fmt.Println(err)
			reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
			return
		}
	} else if req.CourseName == "" && req.CourseCredit != 0 {
		courses, err = db.SelectByCredit(req.CourseCredit, req.Page, req.Size)
		if err != nil {
			fmt.Println(err)
			reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
			return
		}
	} else {
		courses, err = db.SelectCourseByLimit(uint(req.CourseNumberLimit), req.Page, req.Size)
		if err != nil {
			fmt.Println(err)
			reply.Reply(c, http.StatusInternalServerError, "500", err.Error())
			return
		}
	}
	reply.Reply(c, http.StatusOK, "200", courses)
}
