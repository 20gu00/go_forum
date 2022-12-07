package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_forum/common/setUp/logger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//zap以中间件的方式集成进gin
	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
	)

	//ping
	r.GET("/ping", func(ctx *gin.Context) {
		//ctx.String(http.StatusOK,"ok")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "pong",
		})
	})

	SetupRouter(r)
	return r
}
