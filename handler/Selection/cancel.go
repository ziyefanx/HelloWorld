package Selection

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CancelCourseSelection(c *gin.Context) {
	var req DeleteSelectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request, err := db.CancelCourseSelectionInformation(req.StudentID, req.ClassID)
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, req)
	}
	reply.Reply(c, http.StatusOK, "Delete success", request)
}
