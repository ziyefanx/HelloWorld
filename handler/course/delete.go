package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCourseInfo(c *gin.Context) {
	var req DeleteCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "error", err.Error())
		return
	}
	cou, err := db.DeleteCourseInformation(uint(req.CourseID))
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, cou)
	}
	reply.Reply(c, http.StatusOK, "delete success", cou)
}
