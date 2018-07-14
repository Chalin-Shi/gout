package middlewares

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gout/libs/setting"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		adapter := gormadapter.NewAdapter(setting.DBType, setting.DBLink, true)
		enforcer := casbin.NewEnforcer("conf/authz.conf", adapter)
		authorizer := &BasicAuthorizer{enforcer}

		if !authorizer.CheckPermission(c.Request) {
			authorizer.RequirePermission(c.Writer)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	// maid := c.GetStringMap("Maid")
	// name := maid["User"].Name
	name := "chalin"
	return name
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	user := a.GetUserName(r)
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(user, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
}
