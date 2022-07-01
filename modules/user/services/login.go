/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 06:22:23
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

	// 解密客户端送来的用户密码
	decodedPasswd, err := DecodeClientPassword(ctx, pars.Password)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	if !utils.Encrypt.PasswordVerify(decodedPasswd, info.Password) {
		return infoModel.UserInfo{}, fmt.Errorf("用户名或密码错误")
	}

	return info, err
}

func DecodeClientPassword(ctx context.Context, password []byte) ([]byte, error) {
	// rsa 解密客户端送来的用户密码
	rsaDecodedPasswd, err := utils.Encrypt.RsaPrivateDecrypt(password, []byte(
		`-----BEGIN RSA PRIVATE KEY-----
		MIGsAgEAAiEAybD5WijuTcL4TaOEEJbQmRj9cxox8u7RTuuSMltpNmMCAwEAAQIg
		alvpIqJzCI7IYijYe+cMGMBtCfLr6D3YEGHsyg3mwAECEQDjjKOKNnjLj8ILkJW6
		qtUhAhEA4uist+Ezql1qT+iiWyrXAwIRAJ3WSP4vCJ0Sq6O/98wSgWECEQDWu2JY
		UWJXYzfsjza2GACJAhEA0ACcmT1cShxY8FdfJq1SQw==
		-----END RSA PRIVATE KEY-----`))
	if err != nil {
		return nil, err
	}

	base64DecodedPasswd, err := utils.Encrypt.Base64Decode(rsaDecodedPasswd)
	if err != nil {
		return nil, err
	}

	return base64DecodedPasswd, nil
}
