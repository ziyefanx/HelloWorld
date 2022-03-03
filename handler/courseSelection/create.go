package courseSelection

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCourseSelectionInfo(c *gin.Context) {
	var req *model.StudentCourseRelation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	selection, err := db.InsertCourseSelectionInformation(&model.StudentCourseRelation{
		StudentID: req.StudentID,
		ClassID:   req.ClassID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": selection})
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "create success", "data": selection})
}
