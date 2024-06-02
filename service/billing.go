package service

import (
	"github.com/gin-gonic/gin"
	"pocket-book/comm"
	"pocket-book/dao/mysql"
	"pocket-book/models"
	"strconv"
)

func AddExpenses(ctx *gin.Context) {
	reqData := new(models.ParmaAddExpenses)
	if err := ctx.ShouldBindJSON(reqData); err != nil {
		comm.Logger.Error().Msgf("AddExpenses api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	// 写表
	err := mysql.AddExpenses(reqData, userId)
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, "新增记账支出记录成功！")
}
