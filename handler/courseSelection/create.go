package courseSelection

import (
	db "awesomeProject1/dal"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCourseSelectionInfo(c *gin.Context) {
	var req CreateSelectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	num, Err := db.GetCourseSelectionNum(req.ClassID)
	if Err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": Err})
	}
	limit, ERR := db.GetCourseLimit(req.ClassID)
	if ERR != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": ERR})
	}
	if num < limit {
		selection, err := db.InsertCourseSelectionInformation(&model.StudentCourseRelation{
			StudentID: req.StudentID,
			ClassID:   req.ClassID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err, "data": selection})
		}
		c.JSON(http.StatusOK, gin.H{"status": "200", "massage": "create success", "data": selection})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "message": "full number of people"})
}
