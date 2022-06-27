/*
 * @Author: liziwei01
 * @Date: 2022-04-12 15:35:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:37:29
 * @Description: file content
 */
package data

import (
	"context"

	"github.com/liziwei01/gin-file-download/library/oss"
	"github.com/liziwei01/gin-file-download/modules/upload/constant"
)

func GetImageURL(ctx context.Context, fileName string) (string, error) {
	client, err := oss.GetClient(ctx, constant.SERVICE_CONF_OSS_IDIARY)
	if err != nil {
		return "", err
	}

	url, err := client.GetURL(ctx, HEADER_BUCKET_NAME, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}
