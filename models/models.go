package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Chalin-Shi/gout/libs/setting"
	"github.com/Chalin-Shi/gout/libs/util"
)

var db *gorm.DB

type Model struct {
	ID        int   `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func init() {
	var err error

	db, err = gorm.Open(setting.DBType, setting.DBLink)

	if err != nil {
		fmt.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	// db.SingularTable(true)
	db.AutoMigrate(&User{}, &Group{}, &Post{})
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	var root User
	if err := db.Where(User{Email: "chalinsmith@gmail.com"}).Attrs(User{Username: "root", Password: util.Encrypt("123456", "sha256")}).FirstOrCreate(&root).Error; err != nil {
		fmt.Printf("Should not raise any error, but got %v", err)
	}
	db.DB().SetMaxIdleConns(2000)
	db.DB().SetMaxOpenConns(1000)
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().UnixNano() / 1000000
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	nowTime := time.Now().UnixNano() / 1000000
	if _, ok := scope.Get("gorm:update_at"); !ok {
		scope.SetColumn("UpdatedAt", nowTime)
	}
}
