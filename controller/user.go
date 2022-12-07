package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_forum/common"
	"go_forum/model/param"
	"go_forum/service"
)

func RegisterHandler(ctx *gin.Context) {
	// 获取请求参数 参数校验
	p := new(param.RegisterInput)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		// 断言err是不是validator.ValidationErrors 类型(是不是validator支持的错误类型,能否翻译)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译
		return
	}

	// 业务逻辑
	if err := service.Register(p); err != nil {
		zap.L().Error("注册逻辑处理错误", zap.Error(err))
		if errors.Is(err, common.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}

		ResponseError(c, CodeServerBusy)
		return
	}

	// resp
	ResponseSuccess(c, nil)
}
