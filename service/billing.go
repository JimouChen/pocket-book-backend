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
	ResponseSuccess(ctx, "新增记账记录成功！")
}

func EditExpenses(ctx *gin.Context) {
	reqData := new(models.ParmaEditExpenses)
	if err := ctx.ShouldBindJSON(reqData); err != nil {
		comm.Logger.Error().Msgf("AddExpenses api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	// 写表
	err := mysql.EditExpenses(reqData, userId)
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, "编辑记账记录成功！")
}

func DeleteExpenses(ctx *gin.Context) {
	reqData := new(models.ParamDeleteExpenses)
	if err := ctx.ShouldBindJSON(reqData); err != nil {
		comm.Logger.Error().Msgf("DeleteExpenses api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	err := mysql.DeleteExpenses(reqData.BillId, userId)
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, "删除记账记录成功！")
}

// SearchExpenses 支持查个人所有，和条件模糊查询
func SearchExpenses(ctx *gin.Context) {
	reqData := new(models.ParamSearchExpenses)
	if err := ctx.ShouldBindJSON(reqData); err != nil {
		comm.Logger.Error().Msgf("SearchExpenses api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	// 查表
	err, res := mysql.SearchCommExpenses(reqData, userId)
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, res)
}
