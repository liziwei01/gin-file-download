/*
 * @Author: liziwei01
 * @Date: 2022-04-12 13:51:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:33:36
 * @Description: file content
 */
package data

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/liziwei01/gin-lib/library/redis"
	"github.com/liziwei01/gin-file-download/modules/email/constant"

	"github.com/gogf/gf/util/gconv"
)

func VerifyEmailCode(ctx context.Context, email string, code string) (bool, error) {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return false, err
	}

	key := getEmailKey(email)

	realCode, err := client.Get(ctx, key)
	if err != nil {
		return false, err
	}

	return code == realCode, nil
}

// 获取随机六位验证码
func GetRandomVerificationCode(ctx context.Context, email string) (string, error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano()+gconv.Int64(email))).Int31n(1000000))
	return code, nil
}
