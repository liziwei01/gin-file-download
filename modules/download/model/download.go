/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 19:42:48
 * @Description: file content
 */
package model

type DownloadPars struct {
	Path string `form:"path" json:"path" binding:"required"`
}
