package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourseInfo(c *gin.Context) {
	var req CreateCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cou, err := db.InsertCourseInformation(&model.Course{
		CourseID:          req.CourseID,
		CourseName:        req.CourseName,
		CourseCredit:      req.CourseCredit,
		CourseNumberLimit: req.CourseNumberLimit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, cou))
	}
	c.JSON(http.StatusOK, db.Reply("200", "create success", cou))
}
