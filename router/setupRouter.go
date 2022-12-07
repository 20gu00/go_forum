package router

import (
	"github.com/gin-gonic/gin"

	"go_forum/controller"
)

func SetupRouter(r *gin.Engine) {
	//user模块
	r.Group("/user")

	//用户注册
	r.POST("/register", controller.RegisterHandler)
}
