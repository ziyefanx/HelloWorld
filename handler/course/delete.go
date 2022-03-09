package course

import (
	db "awesomeProject1/dal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCourseInfo(c *gin.Context) {
	var req DeleteCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cou, err := db.DeleteCourseInformation(uint(req.CourseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, cou))
	}
	c.JSON(http.StatusOK, db.Reply("200", "delete success", cou))
}
