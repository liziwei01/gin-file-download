/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:35:16
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/liziwei01/gin-lib/library/env"
	"github.com/liziwei01/gin-lib/library/utils"
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

	if !utils.Encrypt.PasswordVerify(decodedPasswd, []byte(info.Password)) {
		return infoModel.UserInfo{}, fmt.Errorf("用户名或密码错误")
	}

	return info, err
}

func DecodeClientPassword(ctx context.Context, password string) ([]byte, error) {
	priKeyPath := getRSAKeyPath()
	priKey, err := ioutil.ReadFile(priKeyPath)
	if err != nil {
		return nil, err
	}

	// rsa 解密客户端送来的用户密码
	password = strings.Replace(password, " ", "", -1)
	passwd, err := utils.Encrypt.Base64Decode([]byte(password))
	if err != nil {
		return nil, err
	}

	rsaDecodedPasswd, err := utils.Encrypt.RsaPrivateDecrypt(passwd, priKey)
	if err != nil {
		return nil, err
	}

	passwd, err = utils.Encrypt.Base64Decode(rsaDecodedPasswd)
	if err != nil {
		return nil, err
	}

	return passwd, nil
}

func getRSAKeyPath() string {
	return env.ConfDir() + "/pem/rsa_private_key.pem"
}
