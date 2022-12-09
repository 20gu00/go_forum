package router

import (
	"github.com/gin-gonic/gin"

	"go_forum/controller"
)

func SetupRouter(r *gin.Engine) {
	//user模块
	user := r.Group("/user")

	//用户注册
	user.POST("/register", controller.RegisterHandler)
	user.POST("/login", controller.LoginHandler)

	//设置为发布gin.SetMode(gin.ReleaseMode),默认debug模式,终端信息输出 debug test release
}
