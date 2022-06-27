/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-26 00:39:38
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"

	emailData "github.com/liziwei01/gin-file-download/modules/email/data"
	userData "github.com/liziwei01/gin-file-download/modules/user/data"

	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	// 1 检验是否已经发过邮件验证码
	hasSent, err := emailData.EmailSent(ctx, pars.Email)
	if err != nil {
		return err
	}

	if !hasSent {
		return fmt.Errorf("请先验证邮箱")
	}

	// // 2 验证邮箱验证码
	correct, err := emailData.VerifyEmailCode(ctx, pars.Email, pars.VerificationCode)
	if err != nil {
		return err
	}
	if !correct {
		return fmt.Errorf("邮箱验证码错误")
	}

	// 3 检查邮箱是否已经注册过
	hasRegisted, err := userData.HasRegisted(ctx, pars.Email)
	if err != nil {
		return err
	}
	if hasRegisted {
		return fmt.Errorf("邮箱已经注册过")
	}

	// 4 注册
	return userData.Register(ctx, pars)
}
