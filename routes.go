package main

import (
	"demo2/controller"
	"demo2/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VJZCI6NiwiZXhwIjoxNzUzMDExMDQwLCJpYXQiOjE3NTI0MDYyNDAsImlzcyI6ImRlbW8yIiwic3ViIjoidXNlciB0b2tlbiJ9.KOBaq7DKdKOOlTUO_BXIN7oKrsxBqoK7OVgkxMHOlEo
