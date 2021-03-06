package middlewares

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Chalin-Shi/gout/libs/e"
	"github.com/Chalin-Shi/gout/libs/setting"
	"github.com/Chalin-Shi/gout/models"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		adapter := gormadapter.NewAdapter(setting.DBType, setting.DBLink, true)
		enforcer := casbin.NewEnforcer("conf/authz.conf", adapter)
		authorizer := &BasicAuthorizer{enforcer}

		if !authorizer.CheckPermission(c) {
			authorizer.RequirePermission(c)
		}
		c.Set("Enforcer", enforcer)
		c.Next()
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserAuthe(c *gin.Context) string {
	maid := c.GetStringMap("Maid")
	user := maid["User"].(models.User)
	authe := user.Username
	if authe != "root" {
		authe = fmt.Sprintf("u_%d", user.ID)
	}
	return authe
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	authe := a.GetUserAuthe(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	return a.enforcer.Enforce(authe, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"status":  e.PERMISSION_DENIED,
		"message": e.GetMsg(e.PERMISSION_DENIED),
	})
	c.Abort()
	return
}
