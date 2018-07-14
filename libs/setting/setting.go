package setting

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	filename string
	RunMode  string

	DBType string
	DBLink string

	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	OSS          map[string]string
	Mail         map[string]string

	Limit      string
	Offset     string
	Secret     string
	SentryKey  string
	BDOSSecret string

	LicenseRestValid string
)

func init() {
	var err error
	name := "dev.ini"
	if os.Getenv("GIN_MODE") == "release" {
		name = "prod.ini"
	}
	filename = fmt.Sprintf("conf/%s", name)

	Cfg, err = ini.Load(filename)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	Cfg, err = ini.Load("conf/base.ini", filename)
	Cfg.BlockMode = false
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadDB()
	LoadServer()
	LoadOSS()
	LoadMail()
	LoadBDOS()
	LoadLicenseRest()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	log.Printf("RunMode = %s", RunMode)
}

func LoadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	DBType = sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()

	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	port := sec.Key("PORT").MustInt()
	DBLink = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Port = sec.Key("PORT").MustInt()
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt()) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt()) * time.Second
}

func LoadOSS() {
	sec, err := Cfg.GetSection("oss")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	OSS = make(map[string]string)
	OSS["AccessKeyId"] = sec.Key("ACCESS_KEY_ID").String()
	OSS["AccessKeySecret"] = sec.Key("ACCESS_KEY_SECRET").String()
	OSS["Bucket"] = sec.Key("BUCKET").String()
	OSS["Prefix"] = sec.Key("PREFIX").String()
	OSS["Endpoint"] = sec.Key("ENDPOINT").String()
	OSS["BaseURL"] = sec.Key("BASEURL").String()
	OSS["Avatar"] = sec.Key("AVATAR").String()
	OSS["Icon"] = sec.Key("ICON").String()
}

func LoadMail() {
	sec, err := Cfg.GetSection("mail")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Mail = make(map[string]string)
	Mail["apiKey"] = sec.Key("APIKey").String()
	Mail["apiURL"] = sec.Key("APIURL").String()
	Mail["apiUser"] = sec.Key("APIUser").String()
	Mail["from"] = sec.Key("From").String()
	Mail["fromName"] = sec.Key("FromName").String()
}

func LoadBDOS() {
	sec, err := Cfg.GetSection("bdos")
	if err != nil {
		log.Fatalf("Fail to get section 'bdos': %v", err)
	}
	BDOSSecret = sec.Key("SECRET").String()
}

func LoadLicenseRest() {
	sec, err := Cfg.GetSection("license.rest")
	if err != nil {
		log.Fatalf("Fail to get section 'license': %v", err)
	}
	LicenseRestValid = sec.Key("VALID").String()
	log.Printf("LicenseRestValid = %s", LicenseRestValid)
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	Secret = sec.Key("SECRET").String()
	SentryKey = sec.Key("SENTRY_KEY").String()
	Limit = sec.Key("PAGE_SIZE").String()
	Offset = "0"
}
