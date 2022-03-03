package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCourseInfo(c *gin.Context) {
	var req *model.Course
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cou, err := db.DeleteCourseInformation(uint(req.CourseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": cou})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Delete success", "data": cou})
}
