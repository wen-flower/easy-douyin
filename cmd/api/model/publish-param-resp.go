package model

import (
	"mime/multipart"

	"github.com/wen-flower/easy-douyin/kitex_gen/common"
)

// PublishListParam 获取视频发布列表参数
type PublishListParam struct {
	UserId int64 `json:"user_id" query:"user_id" vd:"@:$ > 0; msg:'用户ID参数错误'"` // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

// PublishListResp 获取视频发布列表响应数据
type PublishListResp struct {
	BaseResp
	VideoList []*common.VideoInfo `json:"video_list"`
}

// PublishVideoParam 投稿视频参数
type PublishVideoParam struct {
	Data  multipart.FileHeader `form:"data"`
	Title string               `form:"title" vd:"@:mblen($) >= 2 && mblen($) <= 64; msg:'标题长度在2-64之间'"`
}

// PublishVideoResp 投稿视频响应数据
type PublishVideoResp struct {
	BaseResp
}
