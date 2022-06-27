package controllers

import (
	"github.com/liziwei01/gin-file-download/library/response"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	userService "github.com/liziwei01/gin-file-download/modules/user/services"

	"github.com/gin-gonic/gin"
)

func ModifyPassword(ctx *gin.Context) {
	inputs, hasError := getModifyPasswordPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.ModifyPassword(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getModifyPasswordPars(ctx *gin.Context) (userModel.UserPars, bool) {
	var inputs userModel.UserPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
