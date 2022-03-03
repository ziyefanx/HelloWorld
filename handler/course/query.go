package course

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// happy path 不缩进

func QueryCourseInfo(c *gin.Context) {
	var req model.Course
	var cou *model.Course
	var err error
	if err = c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": err.Error()})
		return
	}
	if req.CourseID != 0 {
		cou, err = db.SelectCourseByID(uint(req.CourseID))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
			return
		}
	} else if req.CourseID == 0 && req.CourseName != "" {
		cou, err = db.SelectCourseByName(req.CourseName)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
			return
		}
	} else {
		cou, err = db.SelectByCredit(req.CourseCredit)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": cou})
}
