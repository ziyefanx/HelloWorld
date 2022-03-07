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
	request, err := db.CancelCourseSelectionInfomation(req.ClassID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": err, "data": req})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Delete success", "data": request})
}
