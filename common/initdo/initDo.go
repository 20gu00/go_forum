package initdo

import (
	"flag"
	"fmt"
	"go_forum/common"
	"go_forum/common/snowflake"

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
		fmt.Printf("读取配置文件失败, err:%v\n", err)
		panic(err)
	}

	//初始化logger
	if err := logger.InitLogger(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("初始化logger失败, err:%v\n", err)
		panic(err)
	}
	defer zap.L().Sync() //写入磁盘

	//初始化mysql连接
	if err := mysql.InitMysql(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("初始化mysql失败, err:%v\n", err)
		panic(err)
	}
	defer mysql.DBClose()

	//初始化redis连接
	if err := redis.InitRedis(config.Conf.RedisConfig); err != nil {
		fmt.Printf("初始化redis失败, err:%v\n", err)
		panic(err)
	}
	defer redis.RDBClose()

	//雪花算法生成分布式uid
	if err := snowflake.InitSnowFlake(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("雪花算法生成uid失败, err:%v\n", err)
		//zap.L().Error("雪花算法生成uid失败")
		panic(err)
	}

	//初始化gin内置支持的校验器(validator)的翻译器(en zh)
	if err := common.InitTrans("zh"); err != nil {
		fmt.Printf("初始化validator翻译器失败, err:%v\n", err)
		return
	}
}
