/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 21:48:57
 * @Description: file content
 */
package services

import (
	"context"

	uploadData "github.com/liziwei01/gin-file-download/modules/upload/data"
	uploadModel "github.com/liziwei01/gin-file-download/modules/upload/model"
)

// 图片上传
func UploadImage(ctx context.Context, pars uploadModel.UploadPars) (string, string, error) {
	// 上传oss
	src, err := uploadData.UploadOSS(ctx, pars.UserID, pars.FileHeader)
	if err != nil {
		return "", "", err
	}

	// get url
	url, err := uploadData.GetImageURL(ctx, src)
	if err != nil {
		return "", "", err
	}

	return src, url, nil
}
