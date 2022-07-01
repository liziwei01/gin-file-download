/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 06:06:12
 * @Description: file content
 */
package model

type LoginPars struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password []byte `form:"password" json:"password" binding:"required"`
}
