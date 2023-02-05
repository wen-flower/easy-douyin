package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// FollowerList 实现了 user.UserService 接口
func (us *UserServiceImpl) FollowerList(ctx context.Context, param *user.FollowerListParam) (resp *user.FollowerListResp, err error) {
	resp = new(user.FollowerListResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	userInfos, err := followerList(ctx, param)
	if err != nil {
		return
	}

	resp.UserList = userInfos

	return
}

func followerList(ctx context.Context, param *user.FollowerListParam) ([]*common.UserInfo, error) {
	followerList, err := db.QueryFollowerList(ctx, param.LookUserId)
	if err != nil {
		return nil, err
	}

	// 查询用户的粉丝的 ID 列表
	userIds := make([]int64, 0, len(followerList))
	for _, follow := range followerList {
		userIds = append(userIds, follow.FollowerUser)
	}

	users, err := db.QueryUser(ctx, userIds)

	// 查询登录用户和查询用户粉丝之间的关系
	followList, err := db.QueryFollow(ctx, param.LoggedUserId, userIds)
	if err != nil {
		return nil, err
	}

	return parseUserInfoList(users, followList), nil
}
