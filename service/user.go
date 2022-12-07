package service

import (
	"go_forum/common/snowflake"
	"go_forum/dao/mysql"
	"go_forum/model"
	"go_forum/model/param"
)

func Register(in *param.RegisterInput) error {
	// 判断用户是否已经存在
	if err := mysql.UserIdExist(in.Username); err != nil {
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
