package db

import (
	"context"
	"errors"

	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
	"gorm.io/gorm"
)

// QueryFollow 查询关注状态
func QueryFollow(ctx context.Context, userId int64, userIds []int64) ([]model.Follow, error) {
	var followList []model.Follow
	e := DB.WithContext(ctx).Select(
		model.FollowFollowedUser,
		model.FollowStatus,
	).Where(
		msql.Eq(model.FollowUID), userId,
	).Where(
		msql.In(model.FollowFollowedUser), userIds,
	).Find(&followList).Error

	return followList, e
}

// QueryFollowList 查询粉丝列表
func QueryFollowList(ctx context.Context, userId int64) ([]model.Follow, error) {
	var followList []model.Follow
	e := DB.WithContext(ctx).Select(
		model.FollowFollowedUser,
		model.FollowStatus,
	).Where(
		msql.Eq(model.FollowUID), userId,
	).Where(
		msql.Eq(model.FollowStatus), 1,
	).Find(&followList).Error

	return followList, e
}

// FollowUser 关注用户
func FollowUser(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var curStatus model.Follow
		var notRecord bool

		// 查询是否存在记录
		err := tx.Select(model.FollowID, model.FollowStatus).Where(msql.Eq(model.FollowUID), userId).Where(msql.Eq(model.FollowFollowedUser), toUserId).First(&curStatus).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notRecord = true
			err = nil
		}
		if err != nil {
			return err
		}

		// 已经关注了直接返回
		if curStatus.Status == 1 {
			return nil
		}

		if notRecord { // 没有记录直接创建
			err = tx.Create(&model.Follow{
				UID:          userId,
				FollowedUser: toUserId,
				Status:       1,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&model.Follower{
				UID:          toUserId,
				FollowerUser: userId,
				Status:       1,
			}).Error
			if err != nil {
				return err
			}
		} else { // 有记录则更新
			tx.Model(&model.Follow{}).Where(msql.Eq(model.FollowID), curStatus.ID).Update(model.FollowStatus, 1)
			tx.Model(&model.Follower{}).Where(msql.Eq(model.FollowerUID), toUserId).Where(msql.Eq(model.FollowerFollowerUser), userId).Update(model.FollowerStatus, 1)
		}

		// 更新用户表的冗余数据
		err = tx.Model(&model.User{}).Where(msql.Eq(model.UserUID), userId).Update(model.UserFollowCount, msql.Inc(model.UserFollowCount)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.User{}).Where(msql.Eq(model.UserUID), toUserId).Update(model.UserFollowerCount, msql.Inc(model.UserFollowerCount)).Error

		return err
	})
}

// CancelFollowUser 取消关注用户
func CancelFollowUser(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var curStatus model.Follow

		// 查询是否存在记录
		err := tx.Select(model.FollowID, model.FollowStatus).Where(msql.Eq(model.FollowUID), userId).Where(msql.Eq(model.FollowFollowedUser), toUserId).First(&curStatus).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在直接返回
			return nil
		}
		if err != nil {
			return err
		}

		// 已经取关了直接返回
		if curStatus.Status == 0 {
			return nil
		}

		// 更新记录
		tx.Model(&model.Follow{}).Where(msql.Eq(model.FollowID), curStatus.ID).Update(model.FollowStatus, 0)
		tx.Model(&model.Follower{}).Where(msql.Eq(model.FollowerUID), toUserId).Where(msql.Eq(model.FollowerFollowerUser), userId).Update(model.FollowerStatus, 0)

		// 更新用户表的冗余数据
		err = tx.Model(&model.User{}).Where(msql.Eq(model.UserUID), userId).Update(model.UserFollowCount, msql.Dec(model.UserFollowCount)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.User{}).Where(msql.Eq(model.UserUID), toUserId).Update(model.UserFollowerCount, msql.Dec(model.UserFollowerCount)).Error

		return err
	})
}
