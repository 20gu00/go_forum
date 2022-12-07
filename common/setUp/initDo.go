package setUp

import (
	"flag"

	"go.uber.org/zap"

	"go_forum/common/setUp/config"
	"go_forum/common/setUp/logger"
	"go_forum/dao/mysql"
	"go_forum/dao/redis"
)

func InitDO() {
	var confFile string                           //""
	flag.StringVar(&confFile, "conf", "", "配置文件") //不设置默认值
	flag.Parse()                                  //一次解析即可 支持-flag xxx   --flag xxx   -flag=xxx   --flag=xxx

	//读取配置文件
	if err := config.ConfRead(confFile); err != nil {
		zap.L().Error("读取配置文件失败")
		panic(err)
	}

	//初始化logger
	if err := logger.InitLogger(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("初始化logger失败")
		panic(err)
	}
	defer zap.L().Sync() //写入磁盘

	//初始化mysql连接
	if err := mysql.InitMysql(config.Conf.MysqlConfig); err != nil {
		zap.L().Error("初始化mysql失败")
		panic(err)
	}
	defer mysql.DBClose()

	//初始化redis连接
	if err := redis.InitRedis(config.Conf.RedisConfig); err != nil {
		zap.L().Error("初始化redis失败")
		panic(err)
	}
	defer redis.RDBClose()
}
