package model

import "github.com/wen-flower/easy-douyin/kitex_gen/common"

// VideoFeedParam 视频流接口参数
type VideoFeedParam struct {
	LatestTime *int64 `json:"latest_time" query:"latest_time"` // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

// VideoFeedResp 视频流接口响应数据
type VideoFeedResp struct {
	BaseResp
	VideoList []*common.VideoInfo `json:"video_list"`
}
