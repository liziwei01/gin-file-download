/*
 * @Author: liziwei01
 * @Date: 2022-04-12 15:32:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:53:16
 * @Description: file content
 */
package model

type ImageURL struct {
	FileName string `form:"file_name" json:"file_name" binding:"required"`
}
