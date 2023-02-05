package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// FavoriteAction 点赞视频接口
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.FavoriteActionParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	err = videorpc.FavoriteVideo(ctx, &video.FavoriteVideoParam{
		LoggedUserId: *utils.GetLoggedInUID(req),
		VideoId:      param.VideoId,
		Action:       param.Action == 1,
	})
	if err != nil {
		return
	}

	var resp = new(model.FavoriteActionResp)
	resp.Ok()

	utils.SendJson(req, resp)
}

// FavoriteList 点赞视频列表接口
// @router /douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.FavoriteListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	videoInfos, err := videorpc.FavoriteList(ctx, &video.FavoriteListParam{
		LoggedUserId: utils.GetLoggedInUID(req),
		LookUserId:   param.UserId,
	})
	if err != nil {
		return
	}

	var resp = new(model.FavoriteListResp)
	resp.Ok()
	resp.VideoList = videoInfos

	utils.SendJson(req, resp)
}
