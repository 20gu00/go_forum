package service

import (
	"go_forum/dao/mysql"
	"go_forum/model"
)

func GetCommunityList() ([]*model.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*model.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
