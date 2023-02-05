package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// PublishList 发布列表接口
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.PublishListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	videoInfos, err := videorpc.PublishList(ctx, &video.PublishListParam{
		LookUserId:   param.UserId,
		LoggedInUser: utils.GetLoggedInUID(req),
	})

	resp := new(model.PublishListResp)
	resp.Ok()
	resp.VideoList = videoInfos

	utils.SendJson(req, resp)
}
