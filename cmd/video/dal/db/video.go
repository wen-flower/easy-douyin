package db

import (
	"context"
	"time"

	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
)

// CreateVideo 插入视频数据
func CreateVideo(ctx context.Context, video *model.Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

// QueryVideoList 查询视频列表
func QueryVideoList(ctx context.Context, uid int64) ([]model.Video, error) {
	var videoList []model.Video
	err := DB.WithContext(ctx).Where(msql.Eq(model.VideoUID), uid).Find(&videoList).Error
	return videoList, err
}

// QueryVideoFeed 查询视频流
func QueryVideoFeed(ctx context.Context, latestTime int64, size int) ([]model.Video, error) {
	var videoList []model.Video
	err := DB.WithContext(ctx).Where(
		msql.Gt(model.VideoCreatedAt), time.UnixMilli(latestTime),
	).Order(
		msql.Asc(model.VideoCreatedAt),
	).Limit(size).Find(&videoList).Error
	return videoList, err
}

// QueryVideo 查询视频
func QueryVideo(ctx context.Context, videoIds []int64) ([]model.Video, error) {
	var resp []model.Video
	err := DB.WithContext(ctx).Where(msql.In(model.VideoVid), videoIds).Find(&resp).Error
	return resp, err
}
