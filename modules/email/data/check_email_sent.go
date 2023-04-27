/*
 * @Author: liziwei01
 * @Date: 2022-04-12 13:27:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:31:06
 * @Description: file content
 */
package data

import (
	"context"
	"fmt"

	"github.com/liziwei01/gin-lib/library/redis"
	"github.com/liziwei01/gin-file-download/modules/email/constant"
)

// 检查 redis 是否存在 key: 即是否已经发送过邮件
func EmailSent(ctx context.Context, email string) (bool, error) {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return false, err
	}

	key := getEmailKey(email)

	return client.Exists(ctx, key)
}

func getEmailKey(email string) string {
	return fmt.Sprintf(constant.CACHED_USER_EMAIL, email)
}
