package user

import (
  "net/http"

  "github.com/astaxie/beego/validation"
  "github.com/gin-gonic/gin"

  "upgrade/backend/libs/e"
  "upgrade/backend/libs/logging"
  "upgrade/backend/libs/setting"
  "upgrade/backend/libs/util"
  "upgrade/backend/models"
)

/**
  * @api {post} /user/register POST_USER_REGISTER
  * @apiName POST_USER_REGISTER
  * @apiGroup User
  * @apiPermission Admin User
  *
  * @apiParam {String} email User unique email.
  * @apiParam {String} [password=123456] User password.
  * @apiParam (Authorization) {String} token Only admin users can post this.
  * @apiParamExample {json} Request-Example:
    {
      "email": "Justin@163.com",
      "company": "LinktimeCloud"
      "post": "Backend Engineer",
      "username": "Justin",
      "mobile": "18875906195",
      "password": "123456"
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of user.
  * @apiSuccess {String} data.email User unique email.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {
        "email": "Justin@123.com"
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func AddUser(c *gin.Context) {
  var user models.User
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   make(map[string]interface{}),
    }
    c.Set("response", response)
  }()

  if err := c.ShouldBindJSON(&user); err != nil {
    return
  }

  valid := validation.Validation{}
  email := user.Email
  valid.Required(email, "email").Message("Email is required")
  password := util.RandPassword(8)
  user.Password = util.RandPassword(8)

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  if models.ExistUserByEmail(email) {
    code = e.RECORD_HAS_EXISTED
    return
  }

  models.AddUser(user)
  code = e.SUCCESS

  params := map[string]string{
    "email": email,
    "html":  fmt.Sprintf("Your password is <b>%s</b>, please protect it.", password),
  }
  go util.SendMail(params)
}
