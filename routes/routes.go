package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pocket-book/service"
	"time"
)

func Init() *gin.Engine {
	r := gin.Default()

	// 使用自定义的 CORS 配置
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"UserId", "username", "Content-Type"}, // 添加需要允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")

	// 用户注册
	v1.POST("user/signup", service.SignUp)
	// 用户登陆
	v1.POST("user/login", service.Login)

	// 分类
	v1.POST("/category", service.AddCategory)
	v1.DELETE("/category", service.DeleteCategory)
	v1.GET("/category", service.SearchCategoryByUsername)
	v1.GET("/all_category", service.SearchAllCategory)
	v1.PUT("/category", service.EditCategoryById)

	// 新增支出
	v1.POST("/billing/expenses", service.AddExpenses)
	v1.POST("/billing/search", service.SearchExpenses)
	v1.PUT("/billing/expenses", service.EditExpenses)

	return r
}
