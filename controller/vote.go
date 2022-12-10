package controller

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_forum/common"
	"go_forum/model/param"
	"go_forum/service"

	"github.com/gin-gonic/gin"
)

// 投票

//type VoteData struct {
//	// UserID 从请求中获取当前的用户
//	PostID    int64 `json:"post_id,string"`   // 贴子id
//	Direction int   `json:"direction,string"` // 赞成票(1)还是反对票(-1)
//}

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(param.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			common.RespErr(c, common.CodeInvalidParam)
			return
		}
		errData := common.RemoveTopStruct(errs.Translate(common.Trans)) // 翻译并去除掉错误提示中的结构体标识
		common.RespMsg(c, common.CodeInvalidParam, errData)
		return
	}
	// 获取当前请求的用户的id
	userID, err := common.GetCurrentUserID(c)
	if err != nil {
		common.RespErr(c, common.CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := service.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		common.RespErr(c, common.CodeServerBusy)
		return
	}

	common.RespOK(c, nil)
}
