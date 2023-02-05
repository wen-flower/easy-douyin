package db

import (
	"context"
	"errors"

	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
	"gorm.io/gorm"
)

// QueryFavorite 查询点赞状态
func QueryFavorite(ctx context.Context, userId int64, videoIds []int64) ([]model.Favorite, error) {
	var favoriteList []model.Favorite
	e := DB.WithContext(ctx).Select(
		model.FavoriteVid,
		model.FavoriteStatus,
	).Where(
		msql.Eq(model.FavoriteUID), userId,
	).Where(
		msql.In(model.FavoriteVid), videoIds,
	).Find(&favoriteList).Error

	return favoriteList, e
}

// QueryFavoriteList 查询点赞列表
func QueryFavoriteList(ctx context.Context, userId int64) ([]model.Favorite, error) {
	var favoriteList []model.Favorite
	e := DB.WithContext(ctx).Select(
		model.FavoriteVid,
		model.FavoriteStatus,
	).Where(
		msql.Eq(model.FavoriteUID), userId,
	).Where(
		msql.Eq(model.FavoriteStatus), 1,
	).Find(&favoriteList).Error

	return favoriteList, e
}

// FavoriteVideo 点赞视频
func FavoriteVideo(ctx context.Context, userId int64, videoId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var curStatus model.Favorite
		var notRecord bool

		// 查询是否存在记录
		err := tx.Select(model.FavoriteID, model.FavoriteStatus).Where(msql.Eq(model.FavoriteUID), userId).Where(msql.Eq(model.FavoriteVid), videoId).First(&curStatus).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notRecord = true
			err = nil
		}
		if err != nil {
			return err
		}

		// 已经点赞了直接返回
		if curStatus.Status == 1 {
			return nil
		}

		if notRecord { // 没有记录直接创建
			err = tx.Create(&model.Favorite{
				UID:    userId,
				Vid:    videoId,
				Status: 1,
			}).Error
			if err != nil {
				return err
			}
		} else { // 有记录则更新
			tx.Model(&model.Favorite{}).Where(msql.Eq(model.FavoriteID), curStatus.ID).Update(model.FavoriteStatus, 1)
		}

		// 更新视频表的冗余数据
		err = tx.Model(&model.Video{}).Where(msql.Eq(model.VideoVid), videoId).Update(model.VideoFavoriteCount, msql.Inc(model.VideoFavoriteCount)).Error

		return err
	})
}

// CancelFavoriteVideo 取消点赞视频
func CancelFavoriteVideo(ctx context.Context, userId int64, videoId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var curStatus model.Favorite

		// 查询是否存在记录
		err := tx.Select(model.FavoriteID, model.FavoriteStatus).Where(msql.Eq(model.FavoriteUID), userId).Where(msql.Eq(model.FavoriteVid), videoId).First(&curStatus).Error
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
		tx.Model(&model.Favorite{}).Where(msql.Eq(model.FavoriteID), curStatus.ID).Update(model.FavoriteStatus, 0)

		// 更新视频表的冗余数据
		err = tx.Model(&model.Video{}).Where(msql.Eq(model.VideoID), videoId).Update(model.VideoFavoriteCount, msql.Dec(model.VideoFavoriteCount)).Error

		return err
	})
}
