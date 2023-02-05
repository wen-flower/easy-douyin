package handler

import (
	"encoding/hex"

	"github.com/wen-flower/easy-douyin/cmd/user/consts"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	user "github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"golang.org/x/crypto/sha3"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// 检查 UserServiceImpl 是否实现了 user.UserService 接口
var _ user.UserService = (*UserServiceImpl)(nil)

// 提取出错误的处理流程
func errProcess(baseResp **common.BaseResp, err *error) {
	var resp common.BaseResp
	if *err != nil {
		e := errno.ConvertErr(*err)
		resp.Msg = e.Msg()
		resp.Code = e.Code()
		*err = nil
	} else {
		resp.Code = errno.Success.Code()
		resp.Msg = errno.Success.Msg()
	}
	*baseResp = &resp
}

// 对密码进行摘要
func passwordDigest(password string) string {
	digest := sha3.Sum512([]byte(password + consts.PasswordSecret))
	return hex.EncodeToString(digest[:])
}

// 将 model.User 列表和 model.Follow 列表组合为 common.UserInfo 列表
func parseUserInfoList(users []model.User, followList []model.Follow) []*common.UserInfo {
	resp := make([]*common.UserInfo, 0, len(users))
	followMap := modelFollowToMap(followList)
	for _, _user := range users {
		resp = append(resp, &common.UserInfo{
			FollowCount:   _user.FollowerCount,
			FollowerCount: _user.FollowerCount,
			Id:            _user.UID,
			Followed:      followMap[_user.UID],
			Name:          _user.Username,
		})
	}
	return resp
}

// 将 model.Follow 列表转为 map
func modelFollowToMap(followList []model.Follow) map[int64]bool {
	resp := make(map[int64]bool, len(followList))
	for _, follow := range followList {
		resp[follow.FollowedUser] = follow.Status == 1
	}
	return resp
}
