package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/chat/dal/rdb"
	chat "github.com/wen-flower/easy-douyin/kitex_gen/chat"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
)

// MessageList 实现了 chat.ChatService 接口
func (*ChatServiceImpl) MessageList(ctx context.Context, param *chat.MessageListParam) (resp *chat.MessageListResp, err error) {
	resp = new(chat.MessageListResp)
	defer errProcess(&resp.BaseResp, &err)

	if err = param.IsValid(); err != nil {
		return
	}

	messageInfos, err := messageList(ctx, param)
	if err != nil {
		return
	}

	resp.MessageList = messageInfos

	return
}

func messageList(ctx context.Context, param *chat.MessageListParam) ([]*common.MessageInfo, error) {
	messages, err := rdb.MessageList(ctx, param.LoggedUserId, param.ToUserId)
	if err != nil {
		return nil, err
	}
	return parseMessageInfoList(messages), nil
}
