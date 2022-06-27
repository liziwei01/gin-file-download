package controllers

import (
	"github.com/liziwei01/gin-file-download/library/response"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	userService "github.com/liziwei01/gin-file-download/modules/user/services"

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
