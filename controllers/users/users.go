package users

import (
  "github.com/Unknwon/com"
  "github.com/astaxie/beego/validation"
  "github.com/gin-gonic/gin"

  "gout/libs/e"
  "gout/libs/logging"
  "gout/libs/util"
  "gout/models"
)

/**
  * @api {post} /users POST_USERS
  * @apiName POST_USERS
  * @apiGroup Users
  * @apiPermission Admin User
  *
  * @apiParam {String} email User unique email.
  * @apiParam {String} [password=123456] User password.
  * @apiParam (Authorization) {String} token Only admin users can post this.
  * @apiParamExample {json} Request-Example:
    {
      "email": "Justin@163.com",
      "username": "Justin"
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
  password := user.Password
  valid.Required(email, "email").Message("Email is required")
  valid.Required(password, "password").Message("Password is required")
  user.Password = util.Encrypt(password)

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
}

/**
  * @api {get} /users GET_USERS
  * @apiName GET_USERS
  * @apiGroup Users
  * @apiPermission Authorization User
  *
  * @apiParam {null} null no params.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data.pagination User pagination.
  * @apiSuccess {Number} data.pagination.total User total.
  * @apiSuccess {Number} data.pagination.start User start.
  * @apiSuccess {Number} data.pagination.limit User limit.
  * @apiSuccess {Object[]} data.list User list.
  * @apiSuccess {Number} data.list.id User unique id.
  * @apiSuccess {String} data.list.email User unique email.
  * @apiSuccess {Number} data.list.name User name.
  * @apiSuccess {Timestamp} data.list.createdAt User createdAt.
  * @apiSuccess {Timestamp} data.list.updatedAt User updatedAt.
  * @apiSuccess {Object[]} data.list.group User group.
  * @apiSuccess {String} data.list.group.id Group id.
  * @apiSuccess {String} data.list.group.name Group name.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "pagination": {
          "total": 100,
          "start": 0,
          "limit": 10
        },
        "list": [
          {
            "id": 1,
            "email": "Justin@123.com",
            "name": "Justin",
            "createdAt": 1521113735000,
            "updatedAt": 1521113735000
          },
          {
            "id": 2,
            "email": "Chalin@123.com",
            "name": "Chalin",
            "createdAt": 1521113735000,
            "updatedAt": 1521113735000
          }
        ]
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetUsers(c *gin.Context) {
  var users []models.User
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   users,
    }
    c.Set("response", response)
  }()

  users = models.GetUsers()
  code = e.SUCCESS
}

/**
  * @api {get} /users/:id GET_USERS_UID
  * @apiName GET_USERS_UID
  * @apiGroup Users
  * @apiPermission Authorization User
  *
  * @apiParam {String} id User unique id.
  * @apiParamExample {json} Request-Example:
    {
      "id": 1
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Result of user.
  * @apiSuccess {String} data.id User unique id.
  * @apiSuccess {String} data.email User unique email.
  * @apiSuccess {Number} data.name User name.
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
        "id": 1,
        "email": "Chalin@123.com",
        "name": "Chalin",
        "avatar": "http://bdos-ticket-system.oss-cn-shanghai.aliyuncs.com/avatar.jpg",
        "createdAt": 1521113735000,
        "updatedAt": 1521113735000
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetUserById(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()

  valid := validation.Validation{}
  valid.Min(id, 1, "id").Message("ID must greater than 0")

  code := e.INVALID_PARAMS
  var data interface{}
  if !valid.HasErrors() {
    if models.ExistUserByID(id) {
      data = models.GetUser(id)
      code = e.SUCCESS
    } else {
      code = e.RECORD_NOT_EXIST
    }
  } else {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
  }

  response := map[string]interface{}{
    "status": code,
    "data":   data,
  }
  c.Set("response", response)
}
