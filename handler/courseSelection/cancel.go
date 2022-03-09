package courseSelection

import (
	db "awesomeProject1/dal"
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
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, req))
	}
	c.JSON(http.StatusOK, db.Reply("200", "Delete success", request))
}
