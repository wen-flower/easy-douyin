package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// FollowList 实现了 user.UserService 接口
func (us *UserServiceImpl) FollowList(ctx context.Context, param *user.FollowListParam) (resp *user.FollowListResp, err error) {
	resp = new(user.FollowListResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	userInfos, err := followList(ctx, param)
	if err != nil {
		return
	}

	resp.UserList = userInfos

	return
}

func followList(ctx context.Context, param *user.FollowListParam) ([]*common.UserInfo, error) {
	followList, err := db.QueryFollowList(ctx, param.LookUserId)
	if err != nil {
		return nil, err
	}

	// 查询用户的粉丝的 ID 关注
	userIds := make([]int64, 0, len(followList))
	for _, follow := range followList {
		userIds = append(userIds, follow.FollowedUser)
	}

	users, err := db.QueryUser(ctx, userIds)

	if param.LoggedUserId == nil {
		followList = []model.Follow{}
	} else {
		// 如果不是查询用户的关注列表，则关注状态需要重新查询
		if *param.LoggedUserId != param.LookUserId {
			// 查询登录用户和查询用户关注的用户之间的关系
			followList, err = db.QueryFollow(ctx, *param.LoggedUserId, userIds)
			if err != nil {
				return nil, err
			}
		}
	}

	return parseUserInfoList(users, followList), nil
}
