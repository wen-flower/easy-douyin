package rdb

import (
	"context"
	"strconv"

	"github.com/wen-flower/easy-douyin/cmd/chat/model"
)

// MessageList 获取聊天记录
func MessageList(ctx context.Context, uid int64, formUID int64) ([]*model.Message, error) {
	key := parseMessageKey(uid, formUID)
	count, err := RDB.LLen(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var resp []*model.Message
	err = RDB.LPopCount(ctx, key, int(count)).ScanSlice(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateMessage 添加聊天记录
func CreateMessage(ctx context.Context, uid int64, formUID int64, message *model.Message) error {
	key := parseMessageKey(uid, formUID)
	return RDB.LPush(ctx, key, message).Err()
}

func parseMessageKey(uid int64, formUID int64) string {
	return "message:" + strconv.FormatInt(uid, 10) + ":" + strconv.FormatInt(formUID, 10)
}
