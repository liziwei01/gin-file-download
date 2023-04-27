/*
 * @Author: liziwei01
 * @Date: 2022-04-12 15:31:28
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:09:52
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-lib/library/response"
	uploadModel "github.com/liziwei01/gin-file-download/modules/upload/model"
	uploadService "github.com/liziwei01/gin-file-download/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func GetImageURL(ctx *gin.Context) {
	inputs, hasError := getImageURLPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	src, err := uploadService.GetImageURL(ctx, inputs.FileName)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, src)
}

func getImageURLPars(ctx *gin.Context) (uploadModel.ImageURL, bool) {
	var inputs uploadModel.ImageURL

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
