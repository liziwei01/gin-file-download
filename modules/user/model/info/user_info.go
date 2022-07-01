/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:13:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 06:09:12
 * @Description: file content
 */
package info

type UserInfoRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type UserInfo struct {
	UserID   int64  `json:"user_id" ddb:"user_id"`
	Email    string `json:"email" ddb:"email"`
	Nickname string `json:"nickname" ddb:"nickname"`
	Profile  string `json:"profile" ddb:"profile"`
	Password []byte `json:"password" ddb:"password"`
}

type UserInfoNonsensitive struct {
	UserID   int64  `json:"user_id" ddb:"user_id"`
	Email    string `json:"email" ddb:"email"`
	Nickname string `json:"nickname" ddb:"nickname"`
	Profile  string `json:"profile" ddb:"profile"`
}
