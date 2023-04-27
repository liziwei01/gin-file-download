/*
 * @Author: liziwei01
 * @Date: 2023-03-30 12:07:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:31
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:52:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:57:14
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-lib/library/response"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	userService "github.com/liziwei01/gin-file-download/modules/user/services"

	"github.com/gin-gonic/gin"
)

func ModifyProfile(ctx *gin.Context) {
	inputs, hasError := getModifyProfilePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.ModifyProfile(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getModifyProfilePars(ctx *gin.Context) (userModel.Profile, bool) {
	var inputs userModel.Profile
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
