package db

import (
	"context"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
)

// QueryFollow 查询关注状态
func QueryFollow(ctx context.Context, userId int64, userIds []int64) ([]model.Follow, error) {
	var followList []model.Follow
	e := DB.WithContext(ctx).Select(
		model.FollowUID,
		model.FollowFollowedUser,
	).Where(
		msql.Eq(model.FollowUID), userId,
	).Where(
		msql.In(model.FollowFollowedUser), userIds,
	).Find(&followList).Error

	return followList, e
}
