/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:54:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:25:25
 * @Description: file content
 */
package data

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"strings"

	"github.com/liziwei01/gin-file-download/library/logit"
	"github.com/liziwei01/gin-file-download/library/oss"
	"github.com/liziwei01/gin-file-download/library/utils"
	"github.com/liziwei01/gin-file-download/modules/upload/constant"
)

const (
	// 资源类型
	HEADER_TYPE_INFO = "Content-Type"
	// 资源桶名
	HEADER_BUCKET_NAME = "idiary-image"
)

/**
 * @description: 按照上传路径将文件上传到OSS
 * @param {string} userID 上传来源 决定oss存储路径
 * @param {*multipart.FileHeader} fileHeader
 * @return {*}
 */
func UploadOSS(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (string, error) {
	// 生成文件存储路径
	filepath := genFilePath(userID)

	// 由路径生成完整文件名（包含绝对路径）
	fileName := genFileName(filepath, fileHeader)

	// 获取文件字节流
	byteStream, err := utils.File.GetFileBytes(fileHeader)
	if err != nil {
		err = fmt.Errorf("获取文件字节流失败: %s", err.Error())
		logit.Logger.Error(err)
		return "", err
	}

	// 上传 OSS
	client, err := oss.GetClient(ctx, constant.SERVICE_CONF_OSS_IDIARY)
	if err != nil {
		return "", err
	}

	err = client.Put(ctx, HEADER_BUCKET_NAME, fileName, bytes.NewReader(byteStream))
	if err != nil {
		return "", err
	}

	// 返回文件绝对路径
	return fileName, nil
}

// 生成文件存储路径
func genFilePath(userID string) string {
	return path.Join(userID, "files")
}

// 生成完整文件名（包含绝对路径）
func genFileName(filePath string, fileHeader *multipart.FileHeader) string {
	// 获取文件后缀
	fileSubfix := genFileSubfix(fileHeader.Header.Get(HEADER_TYPE_INFO))

	// 生成唯一文件名
	fileName := utils.UUID.GenUUID() + "." + fileSubfix

	// 返回完整文件名
	return path.Join(filePath, fileName)
}

// 根据 fileType 获取文件后缀  如 'video/mp4'-> 'mp4'
func genFileSubfix(fileType string) string {
	subfixIdx := strings.LastIndex(fileType, "/") + 1
	return fileType[subfixIdx:]
}
