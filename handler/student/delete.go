package student

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteStudentInfo(c *gin.Context) {
	var req *DeleteStudentInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		reply.Reply(c, http.StatusBadRequest, "error", err.Error())
		return
	}
	stu, err := db.DeleteStudentInformation(req.ID)
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, stu)
	}
	reply.Reply(c, http.StatusOK, "Delete success", stu)
}
