package reply

import "github.com/gin-gonic/gin"

func Reply(c *gin.Context, sta int, mes interface{}, info interface{}) {
	c.JSON(sta, gin.H{"status": sta, "message": mes, "data": info})
}
