package controller

import (
	"errors"
	"fmt"
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
			common.RespErr(ctx, common.CodeInvalidParam)
			return
		}

		// err翻译输出
		common.RespMsg(ctx, common.CodeInvalidParam, errs.Translate(common.Trans))
		//common.RespMsg(ctx, common.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
		return
	}

	// 业务逻辑
	if err := service.Register(p); err != nil {
		zap.L().Error("注册逻辑处理错误", zap.Error(err))
		if errors.Is(err, common.ErrorUserExist) {
			common.RespErr(ctx, common.CodeUserExist)
			return
		}

		common.RespErr(ctx, common.CodeServerBusy) //interval
		return
	}

	// resp
	common.RespOK(ctx, nil)
}

func LoginHandler(ctx *gin.Context) {
	// 获取请求参数 参数校验
	p := new(param.LoginInput)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("login 请求参数错误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			common.RespErr(ctx, common.CodeInvalidParam)
			return
		}

		common.RespMsg(ctx, common.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
		return
	}
	// 业务逻辑
	user, err := service.Login(p)
	if err != nil {
		zap.L().Error("login业务逻辑失败", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, common.ErrorUserNotExist) {
			common.RespErr(ctx, common.CodeUserNotExist)
			return
		}
		common.RespErr(ctx, common.CodeInvalidPassword)
		return
	}

	// resp
	common.RespOK(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
