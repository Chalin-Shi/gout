package app

import (
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"upgrade/backend/libs/e"
	"upgrade/backend/libs/logging"
	"upgrade/backend/libs/util"
	"upgrade/backend/models"
)

/**
  * @api {get} /apps/:id/versions GET_APPS_ID_VERSIONS
  * @apiName GET_APPS_ID_VERSIONS
  * @apiGroup Apps
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id App unique id.
  * @apiSuccess {Object[]} data.list App history version list.
  * @apiSuccess {Number} data.list.id App history version id.
  * @apiSuccess {String} data.list.name App history version name.
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
func GetAppVersions(c *gin.Context) {
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

	if !models.ExistAppByID(id) {
		code = e.RECORD_NOT_EXIST
		return
	}

	data["list"] = models.GetVersionsByAppId(id)
	code = e.SUCCESS
}

/**
  * @api {get} /apps/:id/versions/:name GET_APPS_ID_VERSIONS_NAME
  * @apiName GET_APPS_ID_VERSIONS_NAME
  * @apiGroup Apps
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id App unique id.
  * @apiSuccess {String} data.name App name.
  * @apiSuccess {Object} data.version App version.
  * @apiSuccess {String} data.version.current App current version.
  * @apiSuccess {String} data.version.latest App latest version.
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
        "appDesc": "Redis is an in-memory database open-source software project sponsored by Redis Labs.\n It is networked, in-memory, and stores keys with optional durability.",
        "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
        "updatedAt": 1526977135
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func GetAppVersion(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Param("name")

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
	valid.Min(id, 1, "id").Message("ID must greater than 0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		return
	}

	if !models.ExistAppByID(id) {
		code = e.RECORD_NOT_EXIST
		return
	}

	chan1 := make(chan models.App)
	chan2 := make(chan models.Version)
	go func() {
		chan1 <- models.GetApp(id)
	}()
	go func() {
		chan2 <- models.GetVersionByAppId(id, name)
	}()
	app := <-chan1
	version := <-chan2
	data["appId"] = app.ID
	data["appName"] = app.Name
	data["appIcon"] = app.Icon
	data["id"] = version.ID
	data["name"] = version.Name
	data["status"] = version.Status
	data["link"] = version.Link
	data["desc"] = version.Desc
	data["updatedAt"] = version.UpdatedAt
	appDesc := version.AppDesc
	if appDesc == "" {
		appDesc = app.Desc
	}
	data["appDesc"] = appDesc
	code = e.SUCCESS
}

/**
  * @api {post} /apps/:id/versions POST_APPS_ID_VERSIONS
  * @apiName POST_APPS_ID_VERSIONS
  * @apiGroup Apps
  * @apiPermission Authorization User
  *
  * @apiParam {String} name App version name.
  * @apiParam {String} link App sourcelink.
  * @apiParam {String} desc App version desc.
  * @apiParam {String} appDesc App desc.
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {
      "name": "1.0.0",
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "appDesc": "It is networked, in-memory, and stores keys with optional durability."
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id App version unique id.
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
func AddAppVersion(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	defer func() {
		response := map[string]interface{}{
			"status": code,
			"data":   map[string]int{"id": id},
		}
		c.Set("response", response)
	}()

	var version models.Version
	if err := c.ShouldBindJSON(&version); err != nil {
		return
	}

	name := version.Name
	app := map[string]string{"version": name}
	valid := validation.Validation{}
	valid.Required(name, "name").Message("Name is required")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		return
	}

	if models.ExistVersionByAppId(id, name) {
		code = e.RECORD_HAS_EXISTED
		return
	}

	data := make(map[string]interface{})
	data["name"] = version.Name
	data["link"] = version.Link
	data["desc"] = version.Desc
	data["appDesc"] = version.AppDesc
	go models.AddVersionByAppId(id, data)
	go models.EditApp(id, app)
	code = e.SUCCESS

	if code == e.SUCCESS {
		meta := map[string]int64{
			"updatedAt": time.Now().UnixNano() / 1000000,
		}
		message := map[string]interface{}{
			"type": "upgrade",
			"data": meta,
		}
		client := util.Client()
		client.Emit("update", map[string]interface{}{
			"room":    "default",
			"message": message,
		})
	}
}

/**
  * @api {put} /apps/:id/versions/:name PUT_APPS_ID_VERSIONS_NAME
  * @apiName PUT_APPS_ID_VERSIONS_NAME
  * @apiGroup Apps
  * @apiPermission Authorization User
  *
  * @apiParam {String} name App version name.
  * @apiParam {String} link App sourcelink.
  * @apiParam {String} desc App version desc.
  * @apiParam {String} appDesc App desc.
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {
      "status": 0,
      "link": "http://192.168.1.2:8000/linktime-mysql-1.3.0.tar.gz",
      "desc": "It is networked, in-memory, and stores keys with optional durability.",
      "appDesc": "It is networked, in-memory, and stores keys with optional durability."
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id App version unique id.
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
func PutAppVersion(c *gin.Context) {
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

	if !models.ExistVersionByAppId(id, name) {
		code = e.RECORD_NOT_EXIST
		return
	}

	var version models.Version
	if err := c.ShouldBindJSON(&version); err != nil {
		return
	}

	models.EditVersionByAppId(id, version)
	code = e.SUCCESS
}

/**
  * @api {delete} /apps/:id/versions/:name DELETE_APPS_ID_VERSIONS_NAME
  * @apiName DELETE_APPS_ID_VERSIONS_NAME
  * @apiGroup Apps
  * @apiPermission Authorization User
  *
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {}
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {String} data.id App unique id.
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
func DeleteAppVersion(c *gin.Context) {
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

	if !models.ExistVersionByAppId(id, name) {
		code = e.RECORD_NOT_EXIST
		return
	}

	models.DeleteVersionByAppId(id, name)
	code = e.SUCCESS
}
