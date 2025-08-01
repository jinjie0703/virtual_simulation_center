// internal/database/database.go
package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是一个全局的数据库连接池，单例模式
var DB *gorm.DB

// Connect 函数使用传入的 DSN 字符串初始化数据库连接
func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接建立成功")
}
