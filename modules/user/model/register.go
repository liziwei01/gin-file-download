package model

type RegisterPars struct {
	Email            string `form:"email" json:"email" binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" binding:"required"`
	Username         string `form:"username" json:"username" binding:"required"`
	Password         string `form:"password" json:"password" binding:"required"`
}
