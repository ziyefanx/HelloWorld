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
		c.JSON(http.StatusInternalServerError, db.StatusReply("500", Err))
	}
	limit, ERR := db.GetCourseLimit(req.ClassID)
	if ERR != nil {
		c.JSON(http.StatusInternalServerError, db.StatusReply("500", ERR))
	}
	if num >= limit {
		c.JSON(http.StatusInternalServerError, db.StatusReply("500", "Full of people"))
		return
	}
	selection, err := db.InsertCourseSelectionInformation(&model.StudentCourseRelation{
		StudentID: req.StudentID,
		ClassID:   req.ClassID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, db.Reply("500", err, selection))
	}
	c.JSON(http.StatusOK, db.Reply("200", "Create success", selection))
}
