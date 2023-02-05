package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// CommentVideo 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) CommentVideo(ctx context.Context, param *video.CommentVideoParam) (resp *video.CommentVideoResp, err error) {
	resp = new(video.CommentVideoResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	commentInfo, err := commentVideo(ctx, param)
	if err != nil {
		return
	}

	resp.Comment = commentInfo

	return
}

func commentVideo(ctx context.Context, param *video.CommentVideoParam) (*common.CommentInfo, error) {
	comment := &model.Comment{
		Vid:     param.VideoId,
		UID:     param.LoggedUserId,
		Content: param.CommentText,
	}
	err := db.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}
	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: &param.LoggedUserId,
		UserIds:      []int64{param.LoggedUserId},
	})
	if err != nil {
		return nil, err
	}
	if len(userInfos) != 1 {
		return nil, errno.UserNotExistsErr
	}
	return &common.CommentInfo{
		Content:    param.CommentText,
		CreateDate: comment.CreatedAt.Format("01-02"),
		Id:         comment.ID,
		User:       userInfos[0],
	}, nil
}
