package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gout/libs/e"
)

func Formatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		response := c.GetStringMap("response")
		if len(response) == 0 {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		status := response["status"].(string)
		if status == "" {
			status = e.UNKNOW_ERROR
		}
		data := response["data"]
		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"data":    data,
			"message": e.GetMsg(status),
		})
	}
}
