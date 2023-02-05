package model

import "github.com/wen-flower/easy-douyin/kitex_gen/common"

// CommentActionParam 评论视频参数
type CommentActionParam struct {
	// VideoId 视频id
	VideoId int64 `json:"video_id" query:"video_id" vd:"@:$ > 0; msg:'视频ID参数错误'"`
	// Action 操作：1-发布评论，2-删除评论
	Action int16 `json:"action_type" query:"action_type" vd:"@:$ >= 1 && $<=2; msg:'不支持的操作'"`
	// CommentText 用户填写的评论内容，在action_type=1的时候使用
	CommentText *string `json:"comment_text" query:"comment_text" vd:"@:mblen($) >= 2 && mblen($) <= 512; msg:'评论内容长度在2-512之间'"`
	// CommentId 要删除的评论id，在action_type=2的时候使用
	CommentId *int64 `json:"comment_id" query:"comment_id"`
}

// CommentActionResp 评论视频响应数据
type CommentActionResp struct {
	BaseResp
	Comment *common.CommentInfo `json:"comment"`
}

// CommentListParam 获取视频列表参数
type CommentListParam struct {
	// VideoId 视频id
	VideoId int64 `json:"video_id" query:"video_id" vd:"@:$ > 0; msg:'视频ID参数错误'"`
}

// CommentListResp 获取评论列表响应数据
type CommentListResp struct {
	BaseResp
	CommentList []*common.CommentInfo `json:"comment_list"`
}
