package posts

import (
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gout/libs/e"
	"gout/libs/logging"
	"gout/libs/util"
	"gout/models"
)

/**
  * @api {get} /users/:userId/posts GET_USERS_USERID_POSTS
  * @apiName GET_USERS_USERID_POSTS
  * @apiGroup Posts
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
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
func GetUserPosts(c *gin.Context) {
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

	if !models.ExistUserByID(id) {
		code = e.RECORD_NOT_EXIST
		return
	}

	data["list"] = models.GetPostsByUserId(id)
	code = e.SUCCESS
}

/**
  * @api {get} /users/:userId/posts/:id GET_USERS_USERID_POSTS_ID
  * @apiName GET_USERS_USERID_POSTS_ID
  * @apiGroup Posts
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
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
        "userDesc": "Redis is an in-memory database open-source software project sponsored by Redis Labs.\n It is networked, in-memory, and stores keys with optional durability.",
        "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
        "updatedAt": 1526977135
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetUserPost(c *gin.Context) {
	userId := com.StrTo(c.Param("userId")).MustInt()
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
	valid.Min(userId, 1, "userId").Message("ID must greater than 0")
	valid.Min(id, 1, "id").Message("ID must greater than 0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		return
	}

	if !models.ExistUserByID(id) {
		code = e.RECORD_NOT_EXIST
		return
	}

	chan1 := make(chan models.User)
	chan2 := make(chan models.Post)
	go func() {
		chan1 <- models.GetUser(userId)
	}()
	go func() {
		chan2 <- models.GetPostByUserId(userId, id)
	}()
	user := <-chan1
	post := <-chan2
	data["userId"] = user.ID
	data["username"] = user.Username
	data["id"] = post.ID
	data["title"] = post.Name
	data["content"] = post.Content
	data["desc"] = post.Desc
	data["updatedAt"] = post.UpdatedAt
	code = e.SUCCESS
}

/**
  * @api {post} /users/:userId/posts POST_USERS_USERID_POSTS
  * @apiName POST_USERS_USERID_POSTS
  * @apiGroup Posts
  * @apiPermission Authorization User
  *
  * @apiParam {String} name User post name.
  * @apiParam {String} link User sourcelink.
  * @apiParam {String} desc User post desc.
  * @apiParam {String} userDesc User desc.
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {
      "name": "1.0.0",
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "userDesc": "It is networked, in-memory, and stores keys with optional durability."
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
func AddUserPost(c *gin.Context) {
	userId := com.StrTo(c.Param("userId")).MustInt()
	code := e.INVALID_PARAMS

	defer func() {
		response := map[string]interface{}{
			"status": code,
			"data":   map[string]int{"userId": userId},
		}
		c.Set("response", response)
	}()

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return
	}

	title := post.Title
	content := post.Content
	post.UserId = userId
	valid := validation.Validation{}
	valid.Required(title, "title").Message("Title is required")
	valid.Required(content, "content").Message("Content is required")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		return
	}

	if models.ExistPostByUserId(userId, title) {
		code = e.RECORD_HAS_EXISTED
		return
	}

	models.AddPost(post)
	code = e.SUCCESS
}

/**
  * @api {put} /users/:userId/posts/:id PUT_USERS_USERID_POSTS_ID
  * @apiName PUT_USERS_USERID_POSTS_ID
  * @apiGroup Posts
  * @apiPermission Authorization User
  *
  * @apiParam {String} name User post name.
  * @apiParam {String} link User sourcelink.
  * @apiParam {String} desc User post desc.
  * @apiParam {String} userDesc User desc.
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {
      "status": 0,
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "userDesc": "It is networked, in-memory, and stores keys with optional durability."
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
func EditPost(c *gin.Context) {
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

	if !models.ExistPostByUserId(id, name) {
		code = e.RECORD_NOT_EXIST
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return
	}

	models.EditPostByUserId(id, post)
	code = e.SUCCESS
}

/**
  * @api {delete} /users/:userId/posts/:id DELETE_USERS_USERID_POSTS_ID
  * @apiName DELETE_USERS_USERID_POSTS_ID
  * @apiGroup Posts
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
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
func DeletePost(c *gin.Context) {
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

	if !models.ExistPostByUserId(id, name) {
		code = e.RECORD_NOT_EXIST
		return
	}

	models.DeletePostByUserId(id, name)
	code = e.SUCCESS
}
