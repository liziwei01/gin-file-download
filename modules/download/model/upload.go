/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 14:53:43
 * @Description: file content
 */
package model

import "mime/multipart"

type UploadPars struct {
	UserID     string                `form:"user_id" json:"user_id" binding:"required"`
	FileHeader *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}
