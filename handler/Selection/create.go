package Selection

import (
	db "awesomeProject1/dal"
	"awesomeProject1/handler/reply"
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
		reply.Reply(c, http.StatusInternalServerError, "500", Err)
	}
	limit, ERR := db.GetCourseLimit(req.ClassID)
	if ERR != nil {
		reply.Reply(c, http.StatusInternalServerError, "500", ERR)
	}
	if num >= limit {
		reply.Reply(c, http.StatusInternalServerError, "500", "Full of people")
		return
	}
	selection, err := db.InsertCourseSelectionInformation(&model.StudentCourseRelation{
		StudentID: req.StudentID,
		ClassID:   req.ClassID,
	})
	if err != nil {
		reply.Reply(c, http.StatusInternalServerError, err, selection)
	}
	reply.Reply(c, http.StatusOK, "Create success", selection)
}
