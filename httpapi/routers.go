/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-27 12:14:36
 * @Description: 路由分发
 */

package httpapi

import (
	"net/http"

	"github.com/liziwei01/gin-lib/middleware"
	downloadRouters "github.com/liziwei01/gin-file-download/modules/download/routers"
	emailRouters "github.com/liziwei01/gin-file-download/modules/email/routers"
	uploadRouters "github.com/liziwei01/gin-file-download/modules/upload/routers"
	userRouters "github.com/liziwei01/gin-file-download/modules/user/routers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters(handler *gin.Engine) {
	//暂时解决跨域问题
	handler.Use(middleware.CrossRegionMiddleware())
	// router.Use(middleware.CheckTokenMiddleware(), middleware.GetFrequencyControlMiddleware(), middleware.PostFrequencyControlMiddleware(), middleware.MailFrequencyControlMiddleware())
	// init routers
	router := handler.Group("")
	userRouters.Init(router)
	emailRouters.Init(router)
	uploadRouters.Init(router)
	downloadRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iDownloader. Welcome to our offical website!")
	})
}
