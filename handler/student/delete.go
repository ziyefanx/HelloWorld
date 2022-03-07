package student

import (
	db "awesomeProject1/dal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteStudentInfo(c *gin.Context) {
	var req *DeleteStudentInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stu, err := db.DeleteStudentInformation(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": stu})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Delete success", "data": stu})
}
