package user

import (
  "net/http"

  "github.com/astaxie/beego/validation"
  "github.com/gin-gonic/gin"

  "gout/libs/e"
  "gout/libs/logging"
  "gout/libs/util"
  "gout/models"
)

/**
  * @api {post} /auth/login POST_AUTH_LOGIN
  * @apiName POST_AUTH_LOGIN
  * @apiGroup Auth
  * @apiPermission None
  *
  * @apiParam {String} email user email.
  * @apiParam {String} password user password.
  *
  * @apiParamExample {json} Request-Example:
    {
      "email": "admin@linktime.cloud",
      "password": "123456"
    }
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.token Access token.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {
        "token": "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpZCI6IjU4ZGNiY2ViNTFlNGViMGU2ZGNmZTJjMCIsImVtYWlsIjoiTHlubkAxMjMuY29tIiwiYWdlIjoyNSwiZXhwIjoxNDkxNDcwMjczfQ.eAiOMulK26_bLPpZz1sbHPinl_N9-Cb7mEh6LV1a9oBINj8NUcXF8zu4_R5sHACgV79rbnRJNCp9Sdl9-vIXKQ"
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func AuthUser(c *gin.Context) {
  var user models.User
  code := e.INVALID_PARAMS
  data := make(map[string]interface{})

  defer func() {
    c.JSON(http.StatusOK, gin.H{
      "status":  code,
      "data":    data,
      "message": e.GetMsg(code),
    })
  }()

  if err := c.ShouldBindJSON(&user); err != nil {
    return
  }

  valid := validation.Validation{}
  email := user.Email
  password := user.Password
  valid.Required(email, "email").Message("Email is required")
  valid.Required(password, "password").Message("Password is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  id := models.CheckUser(email, util.Encrypt(password))
  if id == 0 {
    code = e.PASSWORD_NOT_MATCH
    return
  }

  token, err := util.GenerateToken(id)
  if err != nil {
    code = e.ERROR_AUTH_TOKEN
    return
  }

  data["token"] = token
  code = e.SUCCESS
}

/**
  * @api {get} /user GET_USER
  * @apiName GET_USER
  * @apiGroup User
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of user.
  * @apiSuccess {String} data.uid User unique id.
  * @apiSuccess {String} data.email User unique email.
  * @apiSuccess {String} data.name User name.
  * @apiSuccess {String} data.avatar User avatar.
  * @apiSuccess {Timestamp} data.createdAt User createdAt.
  * @apiSuccess {Timestamp} data.updatedAt User updatedAt.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {
        "id": "58dcc17b4f959e0f61c056b9",
        "email": "Justin@123.com",
        "name": "Justin",
        "avatar": "http://bdos-ticket-system.oss-cn-shanghai.aliyuncs.com/avatar.jpg",
        "createdAt": 1516849721,
        "updatedAt": 1516849721
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetUser(c *gin.Context) {
  maid := c.GetStringMap("Maid")
  data := maid["User"]

  response := map[string]interface{}{
    "status": e.SUCCESS,
    "data":   data,
  }
  c.Set("response", response)
}

/**
  * @api {put} /user/password PUT_USER_PASSWORD
  * @apiName PUT_USER_PASSWORD
  * @apiGroup User
  *
  * @apiParam (Login) {String} token Only logged in users can post this.
  * @apiParam {String} originPassword User old password.
  * @apiParam {String} password User new password.
  * @apiParamExample {json} Request-Example:
    {
      "originPassword": "123456",
      "password": "654321"
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of user.
  * @apiSuccess {String} data.id User unique id.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    HTTP/1.1 200 OK
    {
      "status": "100000",
      "data": {
        "id": 1
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
type Passwd struct {
  Origin   string `json:"originPassword"`
  Password string `json:"password"`
}

func PutUserPassword(c *gin.Context) {
  maid := c.GetStringMap("Maid")
  user := maid["User"].(models.User)
  id := user.ID
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   map[string]int{"id": id},
    }
    c.Set("response", response)
  }()

  var passwd Passwd
  if err := c.ShouldBindJSON(&passwd); err != nil {
    return
  }

  valid := validation.Validation{}
  valid.Required(passwd.Origin, "originPassword").Message("OriginPassword is required")
  valid.Required(passwd.Password, "Password").Message("Password is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  email := user.Email
  ok := models.CheckUser(email, util.Encrypt(passwd.Origin))
  if ok == 0 {
    code = e.ORIGIN_PASSWORD_ERROR
    return
  }

  models.EditUser(id, map[string]string{"password": util.Encrypt(passwd.Password)})
  code = e.SUCCESS
}
