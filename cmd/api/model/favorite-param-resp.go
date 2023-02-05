package model

import (
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
)

// FavoriteActionParam 点赞视频参数
type FavoriteActionParam struct {
	// VideoId 视频id
	VideoId int64 `json:"video_id" query:"video_id" vd:"@:$ > 0; msg:'视频ID参数错误'"`
	// Action 操作，1-点赞，2-取消点赞
	Action int16 `json:"action_type" query:"action_type" vd:"@:$ >= 1 && $<=2; msg:'不支持的操作'"`
}

// FavoriteActionResp 点赞视频响应数据
type FavoriteActionResp struct {
	BaseResp
}

// FavoriteListParam 获取点赞视频列表参数
type FavoriteListParam struct {
	// UserId 用户ID
	UserId int64 `json:"user_id" query:"user_id" vd:"@:$ > 0; msg:'用户ID参数错误'"`
}

// FavoriteListResp 获取点赞视频列表响应数据
type FavoriteListResp struct {
	BaseResp
	VideoList []*common.VideoInfo `json:"video_list"`
}
