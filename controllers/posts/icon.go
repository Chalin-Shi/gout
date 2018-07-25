package posts

import (
	"encoding/json"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gout/libs/e"
	"gout/libs/logging"
	"gout/libs/setting"
	"gout/libs/util"
	"gout/models"
)

/**
  * @api {post} /apps/:id/icon POST_APPS_ID_ICON
  * @apiName POST_APPS_ID_ICON
  * @apiGroup Apps
  * @apiPermission Authorization User
  *
  * @apiParam {String} file File stream.
  * @apiParam (Authorization) {String} token Only admin user can post this.
  * @apiParamExample {json} Request-Example:
    {
      "file": "stream"
    }
  *
  * @apiSuccess {String} status Status code.
  * @apiSuccess {Object} data Data result.
  * @apiSuccess {Number} data.id App unique id.
  * @apiSuccess {String} data.link App icon link.
  * @apiSuccess {Object} message Descrpition within status code.
  * @apiSuccess {String} message.desc Detail descrption.
  *
  * @apiSuccessExample {json} Success-Response:
    {
      "status": "100000",
      "data": {
        "id": 1,
        "link": "http://bdos-ticket-system.oss-cn-shanghai.aliyuncs.com/avatar.jpg""
      },
      "message": {
        "desc": "Success"
      }
    }
  *
*/
func AddAppIcon(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	link := setting.OSS["Icon"]

	defer func() {
		response := map[string]interface{}{
			"status": code,
			"data":   map[string]interface{}{"id": id, "link": link},
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

	file, err := util.PutObject(c)
	if err != nil {
		code = e.FILE_UPLOAD_FAILED
		return
	}
	link = file.Link
	icon, _ := json.Marshal(file)
	data := map[string]interface{}{"icon": icon}

	if !models.ExistAppByID(id) {
		code = e.RECORD_NOT_EXIST
		return
	}

	models.EditApp(id, data)
	code = e.SUCCESS
}
