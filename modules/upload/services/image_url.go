/*
 * @Author: liziwei01
 * @Date: 2022-04-12 15:34:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:38:01
 * @Description: file content
 */
package services

import (
	"context"

	uploadData "github.com/liziwei01/gin-file-download/modules/upload/data"
)

func GetImageURL(ctx context.Context, fileName string) (string, error) {
	return uploadData.GetImageURL(ctx, fileName)
}
