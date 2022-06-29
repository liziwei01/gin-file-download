/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 05:02:10
 * @Description: file content
 */
package model

// import "mime/multipart"

type DownloadPars struct {
	Path string `form:"path" json:"path" binding:"required"`
	// UserID     string                `form:"user_id" json:"user_id" binding:"required"`
	// FileHeader *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}
