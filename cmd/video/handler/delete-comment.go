package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
)

// DeleteComment 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) DeleteComment(ctx context.Context, param *video.DeleteCommentParam) (resp *video.DeleteCommentResp, err error) {
	resp = new(video.DeleteCommentResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	err = deleteComment(ctx, param)
	if err != nil {
		return
	}

	return
}

func deleteComment(ctx context.Context, param *video.DeleteCommentParam) error {
	return db.DeleteComment(ctx, param.LoggedUserId, param.CommentId, param.VideoId)
}
