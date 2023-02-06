package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/chat"
	"github.com/wen-flower/easy-douyin/pkg/rpc/chatrpc"
)

// MessageList 获取聊天记录
// @router /douyin/message/chat [GET]
func MessageList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.MessageListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	messageInfos, err := chatrpc.MessageList(ctx, &chat.MessageListParam{
		LoggedUserId: *utils.GetLoggedInUID(req),
		ToUserId:     param.ToUserId,
	})
	if err != nil {
		return
	}

	resp := new(model.MessageListResp)
	resp.Ok()
	resp.MessageList = messageInfos

	utils.SendJson(req, resp)
}

// MessageAction 发送信息
// @router /douyin/message/action [POST]
func MessageAction(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.MessageActionParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	loggedUserId := *utils.GetLoggedInUID(req)
	if loggedUserId != param.ToUserId {
		err = chatrpc.MessageAction(ctx, &chat.MessageActionParam{
			LoggedUserId: loggedUserId,
			ToUserId:     param.ToUserId,
			ActionType:   param.Action,
			Content:      param.Content,
		})
		if err != nil {
			return
		}
	}

	resp := new(model.MessageListResp)
	resp.Ok()

	utils.SendJson(req, resp)
}
