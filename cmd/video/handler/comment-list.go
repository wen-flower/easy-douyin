package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// CommentList 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) CommentList(ctx context.Context, param *video.CommentListParam) (resp *video.CommentListResp, err error) {
	resp = new(video.CommentListResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	commentInfos, err := commentList(ctx, param)
	if err != nil {
		return
	}

	resp.CommentList = commentInfos

	return
}

func commentList(ctx context.Context, param *video.CommentListParam) ([]*common.CommentInfo, error) {
	commentList, err := db.QueryCommentList(ctx, param.VideoId)
	if err != nil {
		return nil, err
	}
	userIds := modelCommentToUserIdList(commentList)

	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: param.LoggedUserId,
		UserIds:      userIds,
	})
	if err != nil {
		return nil, err
	}

	return parseCommentInfoList(commentList, userInfos), nil
}
