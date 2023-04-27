/*
 * @Author: liziwei01
 * @Date: 2023-03-30 12:07:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:34:52
 * @Description: file content
 */
package dao

import (
	"context"

	"github.com/liziwei01/gin-lib/library/mysql"
	"github.com/liziwei01/gin-file-download/modules/user/constant"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"email":    pars.Email,
		"password": pars.Password,
		"nickname": pars.Username,
	})

	_, err = client.Insert(ctx, tableName, mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
