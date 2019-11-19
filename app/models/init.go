package model

import (
	"github.com/jinzhu/gorm"
	"time"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB
var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gin2?charset=utf8&parseTime=True&loc=Local")
	// Error
	if err != nil {
		panic(err)
	}
	// log sql
	_ = db.LogMode(true)
	// max idle connection
	db.DB().SetMaxIdleConns(50)
	// connection pool max size
	db.DB().SetMaxOpenConns(100)
	// max connection time
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
func migration() {
	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&Book{})
	DB.Model(&Book{}).Related(&Account{})
}
