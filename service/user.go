package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pocket-book/comm"
	"pocket-book/dao/mysql"
	"pocket-book/models"
	"strings"
)

func Trans2cnForSignUp(err string) string {
	msg := ""
	if strings.Contains(err, "required") {
		msg += "请求参数输入不能为空 "
	}
	if strings.Contains(err, "eqfield") {
		msg += "密码和确认密码不一致"
	}

	return msg
}

func CheckLoginValid(err string) string {
	msg := ""
	if strings.Contains(err, "required") {
		msg += "请求参数输入不能为空 "
	}
	return msg
}

// SignUp 用户注册
func SignUp(ctx *gin.Context) {
	// 参数获取和校验
	userMsg := new(models.ParmaRegister)
	if err := ctx.ShouldBindJSON(userMsg); err != nil {
		comm.Logger.Error().Msgf("sign up with invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, Trans2cnForSignUp(err.Error()))
		return
	}
	//判断是否存在该用户
	if err := mysql.CheckUserIsExist(userMsg.Username); err != nil {
		if errors.Is(err, comm.ErrUserExist) {
			ResponseErrWithMsg(ctx, CodeUserExist, " 注册失败")
			return
		}
		ResponseErrWithMsg(ctx, CodeServerBusy, " 注册失败")
		return
	}
	//调dao写表
	u := &models.ParamUser{
		Username: userMsg.Username,
		Password: userMsg.Password,
	}
	if err := mysql.InsertUser(u); err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	//- 返回响应
	ResponseSuccess(ctx, "注册成功")
}

func Login(ctx *gin.Context) {
	userMsg := new(models.ParmaLogin)
	if err := ctx.ShouldBindJSON(userMsg); err != nil {
		comm.Logger.Error().Msgf("login with invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, CheckLoginValid(err.Error()))
		return
	}

	// 校验用户名和密码
	u := &models.ParamUser{
		Username: userMsg.Username,
		Password: userMsg.Password,
	}
	UserId, err := mysql.CheckLogin(u)
	if err != nil {
		ResponseErr(ctx, CodeUserErrLogin)
		return
	}
	ResponseSuccess(ctx, UserId)
}
