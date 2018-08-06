package util

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"

	"gout/libs/setting"
)

type Icon struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func PutObject(c *gin.Context) (*Icon, error) {
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename

	if setting.RunMode != "release" {
		link := fmt.Sprintf("http://192.168.206.134:%d/static/icons/dashboard.icon.png", setting.Port)
		data := &Icon{filename, link}
		return data, nil
	}

	ossConf := setting.OSS
	ossClient, err := oss.New(ossConf["Endpoint"], ossConf["AccessKeyId"], ossConf["AccessKeySecret"])
	if err != nil {
		return nil, err
	}

	bucket, err := ossClient.Bucket(ossConf["Bucket"])
	if err != nil {
		return nil, err
	}

	object := fmt.Sprintf("%s/%s", ossConf["Prefix"], filename)
	if err := bucket.PutObject(object, file); err != nil {
		fmt.Println("oss put err : ", err)
		return nil, err
	}
	link := fmt.Sprintf("%s/%s", ossConf["BaseURL"], filename)

	data := &Icon{filename, link}
	return data, nil
}
