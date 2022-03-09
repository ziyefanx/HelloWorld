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
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, stu))
	}
	c.JSON(http.StatusOK, db.Reply("200", "Delete success", stu))
}
