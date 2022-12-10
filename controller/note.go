package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_forum/common"
	"go_forum/model"
	"go_forum/model/param"
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
	page, size := common.GetPageInfo(c) //offset limit
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

// GetPostListHandler2 帖子列表接口v2
// @Summary 帖子列表接口v2
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	//c.ShouldBindQuery()  form
	// 初始化结构体时指定初始参数
	p := &param.ParamPostList{
		Page:  1,
		Size:  10,
		Order: param.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		common.RespErr(c, common.CodeInvalidParam)
		return
	}
	data, err := service.GetPostListNew(p) // 更新：合二为一
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy)
		return
	}
	common.RespOK(c, data)
	// 返回响应
}

// 根据社区去查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	// 初始化结构体时指定初始参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
//	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 获取数据
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}
