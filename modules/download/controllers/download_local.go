/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 21:48:02
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-file-download/library/response"
	downloadModel "github.com/liziwei01/gin-file-download/modules/download/model"

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
	
	absPath := "../gin-file-download-data/" + ctx.GetString("email") + inputs.Path
	ctx.File(absPath)
}

func getDownloadLocalPars(ctx *gin.Context) (downloadModel.DownloadPars, bool) {
	var inputs downloadModel.DownloadPars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}
