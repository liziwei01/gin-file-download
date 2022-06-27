/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:56:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 16:01:18
 * @Description: file content
 */
package info

import (
	"context"

	infoDao "github.com/liziwei01/gin-file-download/modules/user/dao/info"
)

func ModifyProfile(ctx context.Context, userID int64, profile string) error {
	return infoDao.ModifyProfile(ctx, userID, profile)
}
