package db

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
)

// QueryFollowerList 查询粉丝列表
func QueryFollowerList(ctx context.Context, userId int64) ([]model.Follower, error) {
	var followerList []model.Follower
	e := DB.WithContext(ctx).Select(
		model.FollowerFollowerUser,
		model.FollowerStatus,
	).Where(
		msql.Eq(model.FollowerUID), userId,
	).Where(
		msql.Eq(model.FollowerStatus), 1,
	).Find(&followerList).Error

	return followerList, e
}
