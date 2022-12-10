package router

import (
	"github.com/gin-gonic/gin"
	"go_forum/controller"
)

func SetupRouter(r *gin.Engine) {
	//user模块
	apiV1 := r.Group("/api/v1")

	user := apiV1.Group("/user") // /api/v1
	{
		user.POST("/register", controller.RegisterHandler)
		user.POST("/login", controller.LoginHandler)
	}

	//apiV1.Use(middleware.JWTMiddleware())

	{
		//社区
		apiV1.GET("/community", controller.CommunityHandler)
		apiV1.GET("/community/:id", controller.CommunityDetailHandler) //  /community/:id  /1  uri参数

		// 帖子
		apiV1.POST("/note", controller.CreatePostHandler)
		apiV1.GET("/note/:id", controller.GetPostDetailHandler) //帖子id
		apiV1.GET("/notelist", controller.GetPostListHandler)

	}

	//设置为发布gin.SetMode(gin.ReleaseMode),默认debug模式,终端信息输出 debug test release
}
