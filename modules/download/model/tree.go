/*
 * @Author: liziwei01
 * @Date: 2022-07-02 19:42:48
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 19:42:48
 * @Description: file content
 */
package model

type TreePars struct {
	Path string `form:"path" json:"path" binding:"required"`
}
