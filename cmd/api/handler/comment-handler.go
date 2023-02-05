package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// CommentAction 评论视频接口
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.CommentActionParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	var comment *common.CommentInfo

	if param.Action == 1 && param.CommentText != nil {
		comment, err = videorpc.CommentVideo(ctx, &video.CommentVideoParam{
			LoggedUserId: *utils.GetLoggedInUID(req),
			VideoId:      param.VideoId,
			CommentText:  *param.CommentText,
		})
	} else if param.CommentId != nil {
		err = videorpc.DeleteComment(ctx, &video.DeleteCommentParam{
			LoggedUserId: *utils.GetLoggedInUID(req),
			VideoId:      param.VideoId,
			CommentId:    *param.CommentId,
		})
	} else {
		err = errno.NotSupportFileTypeErr
	}
	if err != nil {
		return
	}

	var resp = new(model.CommentActionResp)
	resp.Ok()
	resp.Comment = comment

	utils.SendJson(req, resp)
}

// CommentList 评论列表接口
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.CommentListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	commentInfos, err := videorpc.CommentList(ctx, &video.CommentListParam{
		LoggedUserId: utils.GetLoggedInUID(req),
		VideoId:      param.VideoId,
	})
	if err != nil {
		return
	}

	var resp = new(model.CommentListResp)
	resp.Ok()
	resp.CommentList = commentInfos

	utils.SendJson(req, resp)
}
