/*
 * @Author: liziwei01
 * @Date: 2022-07-02 19:41:58
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 21:59:44
 * @Description: file content
 */
package controllers

import (
	"os"
	"path/filepath"

	"github.com/liziwei01/gin-file-download/library/response"
	downloadModel "github.com/liziwei01/gin-file-download/modules/download/model"

	// uploadService "github.com/liziwei01/gin-file-download/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func Tree(ctx *gin.Context) {
	inputs, hasError := getTreePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}

	if inputs.Path[0] != '/' {
		inputs.Path = "/" + inputs.Path
	}

	var files []string
	absPath := "../gin-file-download-data/" + ctx.GetString("email") + inputs.Path
	err := filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		response.StdFailed(ctx, err)
		return
	}

	response.StdSuccess(ctx, files)
}

func getTreePars(ctx *gin.Context) (downloadModel.TreePars, bool) {
	var inputs downloadModel.TreePars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}
