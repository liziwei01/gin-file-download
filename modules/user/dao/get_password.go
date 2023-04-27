/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:45
 * @Description: file content
 */
package dao

import (
	"context"
	"fmt"

	"github.com/liziwei01/gin-lib/library/logit"
	"github.com/liziwei01/gin-lib/library/mysql"
	"github.com/liziwei01/gin-file-download/modules/user/constant"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

const (
	//用户隐私数据表
	USER_PRIVATE_INFO_TABLE = "tb_user_private_info"
)

func GetPassword(ctx context.Context, pars userModel.LoginPars) (string, error) {
	var password []string

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return "", err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"email": pars.Email,
	}

	columns := []string{"password"}

	err = client.Query(ctx, tableName, where, columns, &password)
	if err != nil {
		return "", err
	}

	if len(password) != 1 {
		err := fmt.Errorf("len(password) = %d", len(password))
		logit.Logger.Warn(err.Error())
		return "", err
	}

	return password[0], nil // 密码验证
}
