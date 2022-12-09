package controller

import (
	"go_forum/common"
	"go_forum/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区 频道
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("获取community社区列表失败", zap.Error(err)) // 详细的报错日志记录
		common.RespErr(c, common.CodeServerBusy)           // 不轻易把服务端报错暴露给外面
		return
	}
	common.RespOK(c, data)
}

// 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//如果是模糊查询,那么c.Param就不太够用了,自定义参数结构体
	// 获取社区id
	idStr := c.Param("id") // 获取URL参数 注意和路由对应
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.RespErr(c, common.CodeInvalidParam)
		return
	}

	// 2. 根据id获取社区详情
	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	common.RespOK(c, data)
}
