package dao

import (
	"context"

	"github.com/liziwei01/gin-lib/library/mysql"
	"github.com/liziwei01/gin-file-download/modules/user/constant"
	userModel "github.com/liziwei01/gin-file-download/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"email": pars.Email,
	}

	update_password := map[string]interface{}{
		"password": pars.NewPassword,
	}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_, err = client.Update(ctx, tableName, where, update_password)
	if err != nil {
		return err
	}

	return nil
}
