package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pocket-book/comm"
	"pocket-book/dao/mysql"
	"pocket-book/models"
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
			return
		}
		ResponseErrWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	// 写表
	if err := mysql.AddCategory(cateName); err != nil {
		ResponseErrWithMsg(ctx, CodeServerBusy, fmt.Sprintf("添加分类失败：%s", err.Error()))
		return
	}
	comm.Logger.Info().Msgf("添加分类%s成功", ReqData.Name)
	ResponseSuccess(ctx, fmt.Sprintf("添加分类成功：%s", cateName))
}

func DeleteCategory(ctx *gin.Context) {

}
