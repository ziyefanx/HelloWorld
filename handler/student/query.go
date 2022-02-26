package student

import (
	"awesomeProject1/dal/db"
	"awesomeProject1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// happy path 不缩进

func QueryStudentInfo(c *gin.Context) {
	var req *GetStudentReq
	var stu *model.Student
	var err error
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": err.Error()})
		return
	}
	// 如果传入ID，就只用ID查询
	if req.ID != 0 {
		stu, err = db.SelectStudentByID(req.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
		}
	} else {
		stu, err = db.SelectStudentByNameAndSex(req.Name, int(req.Sex))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "massage": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "massage": stu})
}
