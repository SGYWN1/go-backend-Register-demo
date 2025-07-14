package middleware

import (
	"demo2/common"
	"demo2/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "权限不足"})
			c.Abort()
			return
		}

		//通过验证
		userId := claims.Id
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "权限不足"})
			c.Abort()
			return
		}

		//用户存在将user写入上下文
		c.Set("user", user)

		c.Next()

	}

}
