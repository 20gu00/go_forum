package service

import (
	"fmt"
	"github.com/spf13/viper"
	"go_forum/common/jwt"
	"go_forum/common/snowflake"
	"go_forum/dao/mysql"
	"go_forum/dao/redis"
	"go_forum/model"
	"go_forum/model/param"
	"time"
)

func Register(in *param.RegisterInput) error {
	// 判断用户是否已经存在
	if err := mysql.UserIdExist(in.Username); err != nil {
		fmt.Println("用户已经存在")
		return err
	}
	// 生成UID
	userID := snowflake.GenID()

	// User实例,用于入库
	userDao := &model.User{
		UserID:   userID,
		Username: in.Username,
		Password: in.Password,
	}
	// 数据入库
	return mysql.InsertUser(userDao)
}

func Login(p *param.LoginInput) (user *model.User, err error) {
	user = &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 指针
	if err := mysql.Login(user); err != nil {
		return nil, err //""
	}

	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token

	// 存入redis,用来实现同一个账户只能登陆同一台设备
	if err := redis.UserIdToken(string(user.UserID), token, time.Duration(viper.GetInt("auth.jwt_expire"))*time.Minute); err != nil {
		fmt.Println("user_id和token的存储失败", err)
		return nil, err
	}
	return
}
