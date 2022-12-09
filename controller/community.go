package controller

import (
	"go_forum/common"
	"go_forum/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	common.RespOK(c, data)
}

// 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id") // 获取URL参数
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
