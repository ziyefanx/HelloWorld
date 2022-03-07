package courseSelection

import (
	db "awesomeProject1/dal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourseSelectionInfo(c *gin.Context) {
	var req GetSelectionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := db.GetCourseSelectionInformation(req.StudentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "create success", "data": result})
}
