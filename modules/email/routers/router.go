/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 16:16:36
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	emailController "github.com/liziwei01/gin-file-download/modules/email/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	emailGroup := router.Group("/email")
	{
		emailGroup.POST("/verificationCode", emailController.VerificationCode)
	}
}
