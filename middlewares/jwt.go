package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"upgrade/backend/libs/e"
	"upgrade/backend/libs/setting"
	"upgrade/backend/libs/util"
	"upgrade/backend/models"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = make(map[string]interface{})
		code := e.SUCCESS
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered : ", r)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  code,
					"message": e.GetMsg(code),
					"data":    data,
				})
				c.Abort()
				return
			}
		}()

		secret := setting.Secret
		XProxyHeader := c.GetHeader("X-Special-Proxy-Header")
		token := c.GetHeader("Authorization")
		if XProxyHeader != "" {
			secret = setting.BDOSSecret
			token = XProxyHeader
		} else {
			t := strings.Split(token, "Bearer ")
			if len(t) < 2 {
				code = e.ERROR_AUTH
				return
			}
			token = t[1]
		}

		claims, err := util.ParseToken(token, secret)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_EXPIRED
			return
		}

		clusterId := claims.CLUSTERID
		id := claims.ID
		maid := make(map[string]interface{})
		if clusterId != "" {
			status, err := util.CheckClusterId(clusterId)
			if !status || err != nil {
				code = e.VALIDATION_ERROR
				return
			}
		} else if id > 0 {
			if !models.ExistUserByID(id) {
				code = e.RECORD_NOT_EXIST
				return
			}
			user := models.GetUser(id)
			maid["User"] = user
		} else {
			code = e.ERROR_AUTH
			return
		}
		c.Set("Maid", maid)
		if code != e.SUCCESS {
			panic(errors.New(code))
		}

		c.Next()
	}
}
