package model

//数据表 结构体映射
//time等sqlx自己处理
type User struct {
	UserID   int64  `db:"user_id"` //雪花算法64bit
	Username string `db:"username"`
	Password string `db:"password"`
	//不需要数据库处理
	Token string
}
