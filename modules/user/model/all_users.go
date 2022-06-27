package model

type UserListRequestPars struct {
	PageIndex  uint `form:"page_index" json:"page_index"`
	PageLength uint `form:"page_length" json:"page_length"`
}

type FriendDiaryListRequestPars struct {
	UserID     int64 `form:"user_id" json:"user_id" binding:"required"`
	PageIndex  uint  `form:"page_index" json:"page_index"`
	PageLength uint  `form:"page_length" json:"page_length"`
}

type UserInfo struct {
	UserID   int64  `form:"user_id" json:"user_id" ddb:"user_id" binding:"required"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Nickname string `form:"nickname" json:"nickname" ddb:"nickname"`
	Profile  string `form:"profile" json:"profile" ddb:"image_list"`
	DBTime   int64  `form:"db_time" json:"db_time" ddb:"db_time"`
	Address  string `form:"address" json:"address" ddb:"address"`
}
