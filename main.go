package main

import (
	"demo2/common"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	common.InitDB()
	db := common.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("关闭数据库连接错误: %v", err)
		}
	}()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

// InitDB 连接数据库
