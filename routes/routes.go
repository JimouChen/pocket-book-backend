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

	// 分类
	v1.POST("/category", service.AddCategory)
	v1.DELETE("/category", service.DeleteCategory)
	//v1.GET("/category", service.SearchAllCategory)
	v1.PUT("/category", service.EditCategoryById)

	// 新增支出
	v1.POST("/billing/expenses", service.AddExpenses)
	return r
}
