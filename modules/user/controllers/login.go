/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:23
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-lib/library/cookie"
	"github.com/liziwei01/gin-lib/library/response"
	"github.com/liziwei01/gin-file-download/modules/user/constant"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	userService "github.com/liziwei01/gin-file-download/modules/user/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	inputs, hasError := getLoginPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	info, err := userService.Login(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	encoded, err := cookie.Encode(constant.GIN_FILE_DOWNLOAD_COOKIE_NAME, info.Email)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	ctx.SetCookie(constant.GIN_FILE_DOWNLOAD_COOKIE_NAME, encoded, 0, "/", "", false, true)

	response.StdSuccess(ctx, gin.H{
		constant.GIN_FILE_DOWNLOAD_COOKIE_NAME: encoded,
		"user_id":                              info.UserID,
		"email":                                info.Email,
		"nickname":                             info.Nickname,
		"profile":                              info.Profile,
	})
}

func getLoginPars(ctx *gin.Context) (userModel.LoginPars, bool) {
	var inputs userModel.LoginPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
