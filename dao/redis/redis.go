package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_forum/common/setUp/config"
)

var (
	rdb *redis.Client
	Nil = redis.Nil //一种错误,找不到
)

// Init 初始化连接
func InitRedis(cfg *config.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisAddr, cfg.RedisPort),
		Password:     cfg.RedisPassword, // no password set
		DB:           cfg.DB,            // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdle,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RDBClose() {
	_ = rdb.Close()
}
