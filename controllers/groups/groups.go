package groups

import (
  "fmt"

  "github.com/Unknwon/com"
  "github.com/astaxie/beego/validation"
  "github.com/casbin/casbin"
  "github.com/gin-gonic/gin"

  "gout/libs/e"
  "gout/libs/logging"
  // "gout/libs/util"
  "gout/models"
)

/**
  * @api {get} /groups/:id/users GET_GROUPS_ID_USERS
  * @apiName GET_GROUPS_ID_USERS
  * @apiGroup Groups
  *
  * @apiParam (Authorization) {String} token Only admin group can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id User unique id.
  * @apiSuccess {Object[]} data.list User history post list.
  * @apiSuccess {Number} data.list.id User history post id.
  * @apiSuccess {String} data.list.name User history post name.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 2,
        "list": [{
          "id": 1,
          "name": "1.0.0"
         },
         {
          "id": 2,
          "name": "1.1.0"
        }]
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetGroupUsers(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()
  code := e.INVALID_PARAMS
  var data = map[string]interface{}{"id": id}

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   data,
    }
    c.Set("response", response)
  }()

  valid := validation.Validation{}
  valid.Min(id, 1, "id").Message("ID must greater than 0")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  // if !models.ExistGroupByID(id) {
  //   code = e.RECORD_NOT_EXIST
  //   return
  // }

  // data["list"] = models.GetUsersByGroupId(id)
  code = e.SUCCESS
}

/**
  * @api {get} /groups/:groupId/users/:id GET_GROUPS_GROUPID_USERS_ID
  * @apiName GET_GROUPS_GROUPID_USERS_ID
  * @apiGroup Groups
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin group can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id User unique id.
  * @apiSuccess {String} data.name User name.
  * @apiSuccess {Object} data.post User post.
  * @apiSuccess {String} data.post.current User current post.
  * @apiSuccess {String} data.post.latest User latest post.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 1,
        "name": "1.0.0",
        "status": 1,
        "desc": "Redis is an in-memory database open-source software project sponsored by Redis Labs.\n It is networked, in-memory, and stores keys with optional durability.",
        "groupDesc": "Redis is an in-memory database open-source software project sponsored by Redis Labs.\n It is networked, in-memory, and stores keys with optional durability.",
        "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
        "updatedAt": 1526977135
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetGroupUser(c *gin.Context) {
  groupId := com.StrTo(c.Param("groupId")).MustInt()
  id := com.StrTo(c.Param("id")).MustInt()

  var data = make(map[string]interface{})
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   data,
    }
    c.Set("response", response)
  }()

  valid := validation.Validation{}
  valid.Min(groupId, 1, "groupId").Message("ID must greater than 0")
  valid.Min(id, 1, "id").Message("ID must greater than 0")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  if !models.ExistGroupByID(id) {
    code = e.RECORD_NOT_EXIST
    return
  }

  // chan1 := make(chan models.User)
  // chan2 := make(chan models.Group)
  // go func() {
  //   chan1 <- models.GetUser(groupId)
  // }()
  // go func() {
  //   chan2 <- models.GetGroupByUserId(groupId, id)
  // }()
  // group := <-chan1
  // post := <-chan2
  // data["groupId"] = group.ID
  // data["groupname"] = group.Username
  // data["id"] = post.ID
  // data["title"] = post.Name
  // data["content"] = post.Content
  // data["desc"] = post.Desc
  // data["updatedAt"] = post.UpdatedAt
  code = e.SUCCESS
}

/**
  * @api {post} /groups/:groupId/users/:id POST_GROUPS_GROUPID_USERS_ID
  * @apiName POST_GROUPS_GROUPID_USERS_ID
  * @apiGroup Groups
  * @apiPermission Authorization User
  *
  * @apiParam {String} name User post name.
  * @apiParam {String} link User sourcelink.
  * @apiParam {String} desc User post desc.
  * @apiParam {String} groupDesc User desc.
  * @apiParam (Authorization) {String} token Only admin group can post this.
  * @apiParamExample {json} Request-Example:
    {
      "name": "1.0.0",
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "groupDesc": "It is networked, in-memory, and stores keys with optional durability."
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id User post unique id.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 3
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func AddGroupUser(c *gin.Context) {
  groupId := com.StrTo(c.Param("groupId")).MustInt()
  id := com.StrTo(c.Param("id")).MustInt()

  var data = make(map[string]interface{})
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   data,
    }
    c.Set("response", response)
  }()

  valid := validation.Validation{}
  valid.Min(groupId, 1, "groupId").Message("ID must greater than 0")
  valid.Min(id, 1, "id").Message("ID must greater than 0")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  if !models.ExistGroupByID(groupId) {
    code = e.RECORD_NOT_EXIST
    return
  }

  models.EditUser(id, map[string]int{"group_id": groupId})

  var enforcer *casbin.Enforcer
  if en, ok := c.Get("Enforcer"); ok {
    enforcer = en.(*casbin.Enforcer)
  }
  enforcer.AddGroupingPolicy(fmt.Sprintf("u_%d", id), fmt.Sprintf("g_%d", groupId))

  code = e.SUCCESS
}

/**
  * @api {put} /groups/:groupId/users/:id PUT_GROUPS_GROUPID_USERS_ID
  * @apiName PUT_GROUPS_GROUPID_USERS_ID
  * @apiGroup Groups
  * @apiPermission Authorization User
  *
  * @apiParam {String} name User post name.
  * @apiParam {String} link User sourcelink.
  * @apiParam {String} desc User post desc.
  * @apiParam {String} groupDesc User desc.
  * @apiParam (Authorization) {String} token Only admin group can post this.
  * @apiParamExample {json} Request-Example:
    {
      "status": 0,
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "groupDesc": "It is networked, in-memory, and stores keys with optional durability."
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id User post unique id.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 3
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func EditGroup(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   map[string]int{"id": id},
    }
    c.Set("response", response)
  }()

  name := c.Param("name")
  valid := validation.Validation{}
  valid.Required(name, "name").Message("Name is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  // if !models.ExistGroupByUserId(id, name) {
  //   code = e.RECORD_NOT_EXIST
  //   return
  // }

  // var post models.Group
  // if err := c.ShouldBindJSON(&post); err != nil {
  //   return
  // }

  // models.EditGroupByUserId(id, post)
  code = e.SUCCESS
}

/**
  * @api {delete} /groups/:groupId/users/:id DELETE_GROUPS_GROUPID_USERS_ID
  * @apiName DELETE_GROUPS_GROUPID_USERS_ID
  * @apiGroup Groups
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin group can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id User unique id.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 3
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func DeleteGroup(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()
  code := e.INVALID_PARAMS

  defer func() {
    response := map[string]interface{}{
      "status": code,
      "data":   map[string]int{"id": id},
    }
    c.Set("response", response)
  }()

  name := c.Param("name")
  valid := validation.Validation{}
  valid.Required(name, "name").Message("Name is required")

  if valid.HasErrors() {
    for _, err := range valid.Errors {
      logging.Info(err.Key, err.Message)
    }
    return
  }

  // if !models.ExistGroupByUserId(id, name) {
  //   code = e.RECORD_NOT_EXIST
  //   return
  // }

  // models.DeleteGroupByUserId(id, name)
  code = e.SUCCESS
}
