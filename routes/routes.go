package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	//r.POST("/get_message_id", controller.MsgIdController)

	return r
}
