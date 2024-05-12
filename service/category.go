package service

import (
	"errors"
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
	// 检查分类是否存在，存在则提示，不存在则写入
	if err := mysql.CheckCategoryIsExist(cateName); err != nil {
		if errors.Is(err, comm.ErrCategoryExist) {
			ResponseErr(ctx, CodeCategoryExist)
			return
		}
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
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
	ReqData := new(models.ParamCategoryId)
	if err := ctx.ShouldBindJSON(ReqData); err != nil {
		comm.Logger.Error().Msgf("DeleteCategory api invalid param", err.Error())
		ResponseErrWithMsg(ctx, CodeInvalidParams, err.Error())
		return
	}
	username := ctx.Request.Header.Get(comm.StrUserName)
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))

	if err := mysql.DeleteCategoryById(ReqData.Id, userId); err != nil {
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

func SearchCategoryByUserId(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Request.Header.Get(comm.StrUserId))

	err, res := mysql.SearchCategoryByUserId(userId)
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
