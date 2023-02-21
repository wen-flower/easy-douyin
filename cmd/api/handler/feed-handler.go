package handler

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// VideoFeed 视频流接口
// @router /douyin/feed [GET]
func VideoFeed(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.VideoFeedParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	if param.LatestTime == nil {
        latestTime := time.Now().UnixMilli()
        param.LatestTime = &latestTime
	}

	videoInfos, err := videorpc.VideoFeed(ctx, &video.VideoFeedParam{
		LatestTime:   *param.LatestTime,
		LoggedInUser: utils.GetLoggedInUID(req),
	})
	if err != nil {
		return
	}

	resp := new(model.VideoFeedResp)
	resp.Ok()
	resp.VideoList = videoInfos

	utils.SendJson(req, resp)
}
