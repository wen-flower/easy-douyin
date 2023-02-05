package db

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
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
