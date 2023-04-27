/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:35:10
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:12:20
 * @Description: file content
 */
package info

import (
	"context"
	"fmt"

	"github.com/liziwei01/gin-lib/library/utils"
	infoDao "github.com/liziwei01/gin-file-download/modules/user/dao/info"
	infoModel "github.com/liziwei01/gin-file-download/modules/user/model/info"
)

func BatchGetUserInfo(ctx context.Context, userIDs []int64) ([]infoModel.UserInfo, error) {
	userInfos := make([]infoModel.UserInfo, 0)

	for _, userID := range userIDs {
		if userID <= 0 {
			return nil, fmt.Errorf("invalid user id: %d", userID)
		}
		userInfo, err := infoDao.GetUserInfoByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func GetUserInfoByUserID(ctx context.Context, userID int64) (infoModel.UserInfo, error) {
	info, err := infoDao.GetUserInfoByUserID(ctx, userID)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	return info, nil
}

func GetUserInfoByEmail(ctx context.Context, email string) (infoModel.UserInfo, error) {
	if email == "" {
		return infoModel.UserInfo{}, fmt.Errorf("invalid email: %s", email)
	}

	info, err := infoDao.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	if len(info) != 1 {
		return infoModel.UserInfo{}, fmt.Errorf("invalid user info len: %v", len(info))
	}

	info[0].Password, err = utils.Encrypt.Base64DecodeString(info[0].Password)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	return info[0], nil
}
func GetAllUserInfo(ctx context.Context) ([]infoModel.UserInfo, error) {
	infos, err := infoDao.GetAllUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
