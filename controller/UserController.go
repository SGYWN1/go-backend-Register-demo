package controller

import (
	"demo2/common"
	"demo2/model"
	"demo2/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}

	log.Printf(name, telephone, password)

	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	c.JSON(200, gin.H{"msg": "注册成功"})
}

// Login 用户登录
func Login(c *gin.Context) {
	//获取参数
	DB := common.GetDB()

	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证

	if len(telephone) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码错误"})
		return
	}

	//token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	c.JSON(200, gin.H{
		"code":  200,
		"token": token,
		"msg":   "登录成功"})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})

}

// 判断用户是否存在

// 判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
