/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:58:23
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 16:00:13
 * @Description: file content
 */
package info

import (
	"context"

	"github.com/liziwei01/gin-file-download/library/mysql"
	"github.com/liziwei01/gin-file-download/modules/user/constant"
)

func ModifyProfile(ctx context.Context, userID int64, profile string) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	data := make([]map[string]interface{}, 0)
	update := map[string]interface{}{
		"user_id": userID,
		"profile": profile,
	}
	data = append(data, update)

	_, err = client.InsertOnDuplicate(ctx, tableName, data, update)
	if err != nil {
		return err
	}

	return nil
}
