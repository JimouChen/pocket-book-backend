package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pocket-book/service"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")

	// 用户注册
	v1.POST("user/signup", service.SignUp)
	// 用户登陆
	v1.POST("user/login", service.Login)
	return r
}
