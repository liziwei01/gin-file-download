/*
 * @Author: liziwei01
 * @Date: 2023-03-30 12:07:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:40
 * @Description: file content
 */
package controllers

import (
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	userService "github.com/liziwei01/gin-file-download/modules/user/services"
	"github.com/liziwei01/gin-lib/library/response"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	inputs, hasError := getRegisterPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.Register(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getRegisterPars(ctx *gin.Context) (userModel.RegisterPars, bool) {
	var inputs userModel.RegisterPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
