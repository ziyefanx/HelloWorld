package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCourseInfo(c *gin.Context) {
	var req UpdateCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cou, err := db.UpdateCourseInformation(&model.Course{
		CourseID:          req.CourseID,
		CourseName:        req.CourseName,
		CourseCredit:      req.CourseCredit,
		CourseNumberLimit: req.CourseNumberLimit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": cou})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "Update success", "data": cou})
}
