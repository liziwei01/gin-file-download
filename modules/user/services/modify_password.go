package services

import (
	"context"

	userData "github.com/liziwei01/gin-file-download/modules/user/data"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userData.ModifyPassword(ctx, pars)
}
