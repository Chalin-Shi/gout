package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Chalin-Shi/gout/libs/e"
	"github.com/Chalin-Shi/gout/libs/setting"
	"github.com/Chalin-Shi/gout/libs/util"
	"github.com/Chalin-Shi/gout/models"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = make(map[string]interface{})
		code := e.ERROR_AUTH
		defer func() {
			if code != e.SUCCESS {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  code,
					"message": e.GetMsg(code),
					"data":    data,
				})
				c.Abort()
				return
			}
		}()

		token := c.GetHeader("Authorization")
		if token == "" {
			return
		}
		t := strings.Split(token, "Bearer ")
		if len(t) < 2 {
			return
		}

		claims, err := util.ParseToken(t[1], setting.Secret)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_EXPIRED
			return
		}

		id := claims.ID
		if !models.ExistUserByID(id) {
			code = e.RECORD_NOT_EXIST
			return
		}
		user := models.GetUser(id)
		maid := map[string]interface{}{"User": user}
		c.Set("Maid", maid)
		code = e.SUCCESS

		c.Next()
	}
}
