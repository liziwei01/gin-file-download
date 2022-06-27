package data

import (
	"context"

	userDao "github.com/liziwei01/gin-file-download/modules/user/dao"
	infoDao "github.com/liziwei01/gin-file-download/modules/user/dao/info"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func HasRegisted(ctx context.Context, email string) (bool, error) {
	userInfos, err := infoDao.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return len(userInfos) > 0, nil
}

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	return userDao.Register(ctx, pars)
}
