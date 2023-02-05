package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// QueryUser 实现了 user.UserService 接口
func (us *UserServiceImpl) QueryUser(ctx context.Context, param *user.QueryUserParam) (resp *user.QueryUserResp, err error) {
	resp = new(user.QueryUserResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	userInfoList, err := queryUser(ctx, param)
	if err != nil {
		return
	}

	resp.UserList = userInfoList

	return
}

func queryUser(ctx context.Context, param *user.QueryUserParam) ([]*common.UserInfo, error) {
	users, err := db.QueryUser(ctx, param.UserIds)
	if err != nil {
		return nil, err
	}
	var followList []model.Follow
	if param.LoggedUserId != nil {
		followList, err = db.QueryFollow(ctx, *param.LoggedUserId, param.UserIds)
		if err != nil {
			return nil, err
		}
	}
	return parseUserInfoList(users, followList), nil
}
