/*
 * @Author: liziwei01
 * @Date: 2022-04-18 20:12:54
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:19
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-lib/library/response"
	infoModel "github.com/liziwei01/gin-file-download/modules/user/model/info"

	infoService "github.com/liziwei01/gin-file-download/modules/user/services/info"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	inputs, hasError := getInfoPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	info, err := infoService.Info(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	response.StdSuccess(ctx, gin.H{
		"user_id":  info.UserID,
		"nickname": info.Nickname,
		"profile":  info.Profile,
	})
}

func getInfoPars(ctx *gin.Context) (infoModel.UserInfoRequest, bool) {
	var inputs infoModel.UserInfoRequest
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
