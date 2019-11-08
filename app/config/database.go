package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//db.sh config
const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "wwc_relation_ut"
)

// DB
var DB *sql.DB

// Init database
func InitDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@/gin?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	DB = db
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	return DB
}
