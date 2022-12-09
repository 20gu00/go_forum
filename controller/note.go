package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_forum/common"
	"go_forum/model"
	"go_forum/service"
	"strconv"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	//c.ShouldBindJSON()  // validator --> binding tag
	p := new(model.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		common.RespErr(c, common.CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户的ID
	userID, err := common.GetCurrentUserID(c)
	if err != nil {
		common.RespErr(c, common.CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy)
		return
	}

	// 3. 返回响应
	common.RespOK(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")                      //判断 len()==0 没获取到路径参数
	pid, err := strconv.ParseInt(pidStr, 10, 64) //10进制 64bit
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		common.RespErr(c, common.CodeInvalidParam)
		return
	}

	// 2. 根据id取出帖子数据（查数据库）
	data, err := service.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy)
		return
	}
	// 3. 返回响应
	common.RespOK(c, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := common.GetPageInfo(c)
	// 获取数据
	data, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy)
		return
	}
	common.RespOK(c, data)
	// 返回响应
}
