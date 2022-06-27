/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:30:44
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"

	"github.com/liziwei01/gin-file-download/library/logit"
	emailData "github.com/liziwei01/gin-file-download/modules/email/data"
	emailModel "github.com/liziwei01/gin-file-download/modules/email/model"
)

func VerificationCode(ctx context.Context, pars emailModel.EmailPars) error {
	// 检查 redis 是否存在 key: 即是否已经发送过
	exists, err := emailData.EmailSent(ctx, pars.Email)
	if err != nil {
		return err
	}

	// 如果存在, 则报错
	if exists {
		err = fmt.Errorf("该邮箱60s内已发送过验证码")
		logit.Logger.Info(err.Error())
		return err
	}

	// 如果不存在, 发送验证码
	// 获取随机六位验证码
	code, err := emailData.GetRandomVerificationCode(ctx, pars.Email)
	if err != nil {
		return err
	}

	// 发送邮件
	err = emailData.SendEmail(ctx, pars.Email, code)
	if err != nil {
		return err
	}

	// 存储验证码
	err = emailData.SaveVerificationCode(ctx, pars.Email, code)
	if err != nil {
		return err
	}

	return nil
}
