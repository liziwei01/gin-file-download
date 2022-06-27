package model

type UserPars struct {
	Email       string `form:"email" json:"email" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}
