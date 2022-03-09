package db

import "github.com/gin-gonic/gin"

func Reply(sta string, mes interface{}, info interface{}) interface{} {
	return gin.H{"status": sta, "message": mes, "data": info}
}
func StatusReply(sta string, info interface{}) interface{} {
	return gin.H{"status": sta, "message": info}
}
