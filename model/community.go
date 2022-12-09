package model

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // 字段为空则不展示
	CreateTime   time.Time `json:"create_time" db:"create_time"`             //有时候前端希望拿到的是时间戳,那就转换int64
}
