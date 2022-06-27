/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-28 02:00:08
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"

	"github.com/liziwei01/gin-file-download/library/utils"
	infoData "github.com/liziwei01/gin-file-download/modules/user/data/info"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
	infoModel "github.com/liziwei01/gin-file-download/modules/user/model/info"
)

func Login(ctx context.Context, pars userModel.LoginPars) (infoModel.UserInfo, error) {
	info, err := infoData.GetUserInfoByEmail(ctx, pars.Email)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	if !utils.Encrypt.PasswordVerifyString(pars.Password, info.Password) {
		return infoModel.UserInfo{}, fmt.Errorf("用户名或密码错误")
	}

	return info, err
}
