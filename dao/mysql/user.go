package mysql

import (
	"database/sql"
	"go_forum/common"
	"go_forum/model"
)

// 插入一条记录进user表
func InsertUser(user *model.User) (err error) {
	// 对密码进行加密
	user.Password = common.MD5(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//根据id获取用户信息
func GetUserById(uid int64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

// 判断用户是否存在
func UserIdExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		// 数据库get操作的错误
		return err
	}

	if count > 0 {
		return common.ErrorUserExist
	}

	return //nil
}

// 登录
func Login(user *model.User) (err error) {
	//请求参数
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows { //查询不到对应的记录
		return common.ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := common.MD5(oPassword)
	if password != user.Password {
		return common.ErrorInvalidPassword
	}
	return
}
