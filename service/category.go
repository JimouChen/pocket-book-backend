package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pocket-book/comm"
	"pocket-book/dao/mysql"
	"pocket-book/models"
	"strconv"
)

func AddCategory(ctx *gin.Context) {
	ReqData := new(models.ParamCategories)
	if err := ctx.ShouldBindJSON(ReqData); err != nil {
		comm.Logger.Error().Msgf("AddCategory api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	cateName := ReqData.Name
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))

	// 写表
	if err := mysql.AddCategory(cateName, userId); err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, fmt.Sprintf("添加分类失败：%s", err.Error()))
		return
	}
	comm.Logger.Info().Msgf("添加分类%s成功", ReqData.Name)
	ResponseSuccess(ctx, fmt.Sprintf("添加分类成功：%s", cateName))
}

func DeleteCategory(ctx *gin.Context) {
	//categoryId := ctx.Query(comm.StrCategoryId)
	var reqData models.ParamDeleteCategory
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		comm.Logger.Error().Msgf("DelCategory api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	username := ctx.Request.Header.Get(comm.StrUserName)
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	//CategoryId, _ := strconv.Atoi(categoryId)

	if err := mysql.DeleteCategoryByNames(reqData.CategoryNames, userId); err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, fmt.Sprintf("删除分类失败：%s", err.Error()))
		return
	}
	comm.Logger.Info().Msgf("用户 %s 删除了分类!", username)
	ResponseSuccess(ctx, "删除分类成功!")
}

func EditCategoryById(ctx *gin.Context) {
	ReqData := new(models.ParamEditCategory)
	if err := ctx.ShouldBindJSON(ReqData); err != nil {
		comm.Logger.Error().Msgf("EditCategoryById api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	username := ctx.Request.Header.Get(comm.StrUserName)
	if err := mysql.EditCategoryById(ReqData.Id, ReqData.Name); err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, fmt.Sprintf("删除分类失败：%s", err.Error()))
		return
	}
	comm.Logger.Info().Msgf("用户 %s 编辑了分类!", username)
	ResponseSuccess(ctx, "编辑分类成功!")
}

func SearchCategoryByUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	//userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))
	err, res := mysql.SearchCategoryByUsername(username)
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, res)
}

func SearchAllCategory(ctx *gin.Context) {
	err, res := mysql.SearchAllCategory()
	if err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, res)
}
