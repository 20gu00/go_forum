package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_forum/common"
	"go_forum/common/setUp/config"
	"go_forum/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bluebell

// @contact.name liwenzhou
// @contact.url http://www.liwenzhou.com

// @host 127.0.0.1:8084
// @BasePath /api/v1
func main() {
	common.InitDO()
	r := router.InitRouter()

	server := http.Server{
		Addr:           fmt.Sprintf(":%d", config.Conf.Port),
		Handler:        r,
		ReadTimeout:    time.Duration(config.Conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Conf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << config.Conf.MaxHeader,
	}

	go func() {
		zap.L().Info("[Info]",
			zap.String("程序名称", viper.GetString("app_name")),
			zap.String("程序版本", viper.GetString("version")),
			zap.Int("server port", viper.GetInt("app_port")),
		)
		fmt.Println("server port:", viper.GetInt("app_port"))
		if err := server.ListenAndServe(); err != nil { //阻塞
			zap.L().Info("web server启动失败", zap.Error(err))
		}

	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("server不正常退出,shutdown", zap.Error(err))
	}

	zap.L().Info("server退出了")
}
