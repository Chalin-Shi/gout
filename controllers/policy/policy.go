package policy

import (
  "fmt"

  "github.com/astaxie/beego/validation"
  "github.com/casbin/casbin"
  "github.com/gin-gonic/gin"

  "gout/libs/e"
  "gout/libs/logging"
)

type Policy struct {
  ID     int    `json:"id"`
  Path   string `json:"path"`
  Method string `json:"method"`
}

/**
  * @api {post} /policy POST_POLICY
  * @apiName POST_POLICY
  * @apiGroup Policy
  * @apiPermission Admin Policy
  *
  * @apiParam {String} email Policy unique email.
  * @apiParam {String} [password=123456] Policy password.
  * @apiParam (Authorization) {String} token Only admin policy can post this.
  * @apiParamExample {json} Request-Example:
    {
      "id": 1,
      "path": "/api/users/:id",
      "method": "GET"
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of policy.
  * @apiSuccess {String} data.email Policy unique email.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {},
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func AddPolicy(c *gin.Context) {
  var policy Policy
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   make(map[string]interface{}),
    }
    c.Set("response", response)
  }()

  if err := c.ShouldBindJSON(&policy); err != nil {
    return
  }

  valid := validation.Validation{}
  id := policy.ID
  path := policy.Path
  method := policy.Method
  valid.Min(id, 1, "id").Message("ID must greater than 0")
  valid.Required(path, "path").Message("Path is required")
  valid.Required(method, "method").Message("Method is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  var enforcer *casbin.Enforcer
  if en, ok := c.Get("Enforcer"); ok {
    enforcer = en.(*casbin.Enforcer)
  }
  enforcer.AddPolicy(fmt.Sprintf("%d", id), path, method)
  enforcer.SavePolicy()

  code = e.SUCCESS
}

/**
  * @api {delete} /policy DELETE_POLICY
  * @apiName DELETE_POLICY
  * @apiGroup Policy
  * @apiPermission Admin Policy
  *
  * @apiParam {String} email Policy unique email.
  * @apiParam {String} [password=123456] Policy password.
  * @apiParam (Authorization) {String} token Only admin policy can post this.
  * @apiParamExample {json} Request-Example:
    {
      "id": 1,
      "path": "/api/users/:id",
      "method": "GET"
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of policy.
  * @apiSuccess {String} data.email Policy unique email.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {},
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func DelPolicy(c *gin.Context) {
  var policy Policy
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   make(map[string]interface{}),
    }
    c.Set("response", response)
  }()

  if err := c.ShouldBindJSON(&policy); err != nil {
    return
  }

  valid := validation.Validation{}
  id := policy.ID
  path := policy.Path
  method := policy.Method
  valid.Min(id, 1, "id").Message("ID must greater than 0")
  valid.Required(path, "path").Message("Path is required")
  valid.Required(method, "method").Message("Method is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  var enforcer *casbin.Enforcer
  if en, ok := c.Get("Enforcer"); ok {
    enforcer = en.(*casbin.Enforcer)
  }
  enforcer.RemovePolicy(fmt.Sprintf("%d", id), path, method)
  enforcer.SavePolicy()

  code = e.SUCCESS
}
