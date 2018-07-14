package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"gout/libs/setting"
)

func GetPage(c *gin.Context) (int, int) {
	limit := com.StrTo(c.DefaultQuery("limit", setting.Limit)).MustInt()
	offset := com.StrTo(c.DefaultQuery("start", setting.Offset)).MustInt()

	return limit, offset
}
