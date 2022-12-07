package mysql

import (
	"fmt"
	"go_forum/common/setUp/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitMysql(cfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.UserName, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlPort, cfg.DBName)
	//viper.GetString("mysql.user_name"),
	//viper.GetString("mysql.password"),
	//viper.GetString("mysql.addr"),
	//viper.GetInt("mysql.port"),
	//viper.GetString("mysql.db_name"),
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxConnection)
	db.SetMaxIdleConns(cfg.MaxIdle)
	return
}

//包内全局
func DBClose() {
	_ = db.Close()
}
