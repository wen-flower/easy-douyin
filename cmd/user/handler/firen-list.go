package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/mmath"
)

// FriendList 实现了 user.UserService 接口
func (us *UserServiceImpl) FriendList(ctx context.Context, param *user.FriendListParam) (resp *user.FriendListResp, err error) {
	resp = new(user.FriendListResp)
	defer errProcess(&resp.BaseResp, &err)

	if err = param.IsValid(); err != nil {
		return
	}

	userInfos, err := friendList(ctx, param)
	if err != nil {
		return
	}

	resp.UserList = userInfos

	return
}

func friendList(ctx context.Context, param *user.FriendListParam) ([]*common.UserInfo, error) {
	followList, err := db.QueryFollowList(ctx, param.LookUserId)
	if err != nil {
		return nil, err
	}
	followerList, err := db.QueryFollowerList(ctx, param.LookUserId)
	if err != nil {
		return nil, err
	}

	followMap := modelFollowToMap(followList)

	// 查询用户的好友的 ID
	userIds := make([]int64, 0, mmath.MinInt(len(followList), len(followerList)))
	for _, follow := range followerList {
		if _, ok := followMap[follow.FollowerUser]; ok {
			userIds = append(userIds, follow.FollowerUser)
		}
	}

	users, err := db.QueryUser(ctx, userIds)
	if err != nil {
		return nil, err
	}

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
