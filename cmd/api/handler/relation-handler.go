package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// FollowAction 关注操作
// @router /douyin/relation/action [POST]
func FollowAction(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.FollowActionParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	err = userrpc.FollowUser(ctx, &user.FollowUserParam{
		LoggedUserId: *utils.GetLoggedInUID(req),
		FollowUserId: param.ToUserId,
		Action:       param.Action == 1,
	})
	if err != nil {
		return
	}

	var resp = new(model.FollowActionResp)
	resp.Ok()

	utils.SendJson(req, resp)
}

// FollowList 关注列表
// @router /douyin/relation/follow/list [GET]
func FollowList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.FollowListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	userInfos, err := userrpc.FollowList(ctx, &user.FollowListParam{
		LoggedUserId: utils.GetLoggedInUID(req),
		LookUserId:   param.LookUserId,
	})
	if err != nil {
		return
	}

	var resp = new(model.FollowListResp)
	resp.Ok()

	resp.UserList = userInfos

	utils.SendJson(req, resp)
}

// FollowerList 粉丝列表
// @router /douyin/relation/follower/list [GET]
func FollowerList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.FollowerListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	userInfos, err := userrpc.FollowerList(ctx, &user.FollowerListParam{
		LoggedUserId: utils.GetLoggedInUID(req),
		LookUserId:   param.LookUserId,
	})
	if err != nil {
		return
	}

	var resp = new(model.FollowerListResp)
	resp.Ok()

	resp.UserList = userInfos

	utils.SendJson(req, resp)
}
