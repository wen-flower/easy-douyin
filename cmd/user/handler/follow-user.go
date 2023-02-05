package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// FollowUser 实现了 user.UserService 接口
func (us *UserServiceImpl) FollowUser(ctx context.Context, param *user.FollowUserParam) (resp *user.FollowUserResp, err error) {
	resp = new(user.FollowUserResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	err = followUser(ctx, param)

	return
}

func followUser(ctx context.Context, param *user.FollowUserParam) error {
	if param.Action {
		return db.FollowUser(ctx, param.LoggedUserId, param.FollowUserId)
	}
	return db.CancelFollowUser(ctx, param.LoggedUserId, param.FollowUserId)
}
