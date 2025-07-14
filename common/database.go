package common

import (
	"demo2/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "gy123456" // 请换成你真实的密码
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	// 自动建表（只执行一次也可以删掉）
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db

}
