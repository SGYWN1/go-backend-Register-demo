package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;unique"`
	Telephone string `gorm:"type:varchar(110);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
}

func main() {
	db := InitDB()
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
	r.POST("/ping", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		if len(telephone) != 11 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "手机号必须为11位数"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码不能小于6位"})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)

		}

		log.Printf(name, telephone, password)

		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果

		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		return false
	}
	return true

}

// RandomString 随机生成名称
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}

// InitDB 连接数据库
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "gy123456" // 请换成你真实的密码
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	// 自动建表（只执行一次也可以删掉）
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}

	return db
}
