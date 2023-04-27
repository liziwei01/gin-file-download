/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-27 14:37:19
 * @Description: file content
 */
package controllers

import (
	"path/filepath"

	"github.com/liziwei01/gin-file-download/library/file"
	"github.com/liziwei01/gin-file-download/modules/download/constant"
	downloadModel "github.com/liziwei01/gin-file-download/modules/download/model"
	"github.com/liziwei01/gin-lib/library/response"

	// uploadService "github.com/liziwei01/gin-file-download/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func DownloadLocal(ctx *gin.Context) {
	inputs, hasError := getDownloadLocalPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}

	if inputs.Path[0] != '/' {
		inputs.Path = "/" + inputs.Path
	}
	client, err := file.GetClient(ctx, constant.SERVICE_CONF_FILE)
	if err != nil {
		return
	}
	absPath := client.Location() + ctx.GetString("email") + inputs.Path
	rlvPath := filepath.Base(absPath)
	ctx.FileAttachment(absPath, rlvPath)
}

func getDownloadLocalPars(ctx *gin.Context) (downloadModel.DownloadPars, bool) {
	var inputs downloadModel.DownloadPars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}
