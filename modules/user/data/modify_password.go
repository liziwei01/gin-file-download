package data

import (
	"context"

	userDao "github.com/liziwei01/gin-file-download/modules/user/dao"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userDao.ModifyPassword(ctx, pars)
}
