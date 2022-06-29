/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 05:23:14
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/liziwei01/gin-file-download/middleware"
	downloadController "github.com/liziwei01/gin-file-download/modules/download/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	downloadGroup := router.Group("/gin-file-download/api/download")
	downloadGroup.Use(middleware.CheckLoginMiddleware())
	{
		downloadGroup.GET("/local", downloadController.DownloadLocal)
	}
}
