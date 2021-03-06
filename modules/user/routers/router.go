/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 22:02:35
 * @Description: file content
 */
package routers

import (
	userControllers "github.com/liziwei01/gin-file-download/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		// userGroup.POST("/register", userControllers.Register)
		userGroup.GET("/login", userControllers.Login)
		// userGroup.POST("/modifyPassword", userControllers.ModifyPassword)
		// userGroup.POST("/getUserInfo", userControllers.Info)
	}
}
