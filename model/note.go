package model

import "time"

//结构体内存对齐 同类型的字段放一块 节省内存空间(unsafe.Sizeof(Post))
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`                            // 帖子id   解决失真问题 前端->json->反序列化go数据int64  int64->序列化成json string
	AuthorID    int64     `json:"author_id" db:"author_id"`                          // 作者id
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 帖子创建时间
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"` // 作者
	VoteNum          int64              `json:"vote_num"`    // 投票数
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息  //单独放一层
}
