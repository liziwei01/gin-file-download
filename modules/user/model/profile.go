/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:53:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:53:50
 * @Description: file content
 */
package model

type Profile struct {
	UserID      int64  `form:"user_id" json:"user_id" binding:"required"`
	UserProfile string `form:"user_profile" json:"user_profile" binding:"required"`
}
