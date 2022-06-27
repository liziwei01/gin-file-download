/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 21:49:57
 * @Description: file content
 */
package controllers

import (
	"fmt"
	"time"

	"github.com/liziwei01/gin-file-download/library/logit"
	"github.com/liziwei01/gin-file-download/library/response"
	uploadModel "github.com/liziwei01/gin-file-download/modules/upload/model"
	uploadService "github.com/liziwei01/gin-file-download/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) {
	inputs, hasError := getUploadImagePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	src, url, err := uploadService.UploadImage(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"file_name": src,
		"url":       url,
	})
}

func getUploadImagePars(ctx *gin.Context) (uploadModel.UploadPars, bool) {
	var inputs uploadModel.UploadPars

	startFormat := time.Now()
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	costFormat := time.Since(startFormat)

	logit.Logger.Info(fmt.Sprintf("get upload par from request cost=[%s], file bytes is %d", costFormat, inputs.FileHeader.Size))

	return inputs, false
}
