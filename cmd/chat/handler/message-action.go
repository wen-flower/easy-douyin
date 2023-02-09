package handler

import (
	"context"
	"time"

	"github.com/wen-flower/easy-douyin/cmd/chat/dal/rdb"
	"github.com/wen-flower/easy-douyin/cmd/chat/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/chat"
)

// MessageAction 实现了 chat.ChatService 接口
func (*ChatServiceImpl) MessageAction(ctx context.Context, param *chat.MessageActionParam) (resp *chat.MessageActionResp, err error) {
	resp = new(chat.MessageActionResp)
	defer errProcess(&resp.BaseResp, &err)

	if err = param.IsValid(); err != nil {
		return
	}

	err = messageAction(ctx, param)

	return
}

func messageAction(ctx context.Context, param *chat.MessageActionParam) error {
	now := time.Now()
	return rdb.CreateMessage(ctx, param.ToUserId, param.LoggedUserId, &model.Message{
		Id:         now.UnixMicro(),
		Content:    param.Content,
		CreateTime: now,
	})
}
