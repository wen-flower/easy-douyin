package model

import "github.com/wen-flower/easy-douyin/kitex_gen/common"

// MessageListParam 获取聊天记录列表参数
type MessageListParam struct {
	ToUserId int64 `json:"to_user_id" query:"to_user_id"`
}

// MessageListResp 获取聊天记录响应数据
type MessageListResp struct {
	BaseResp
	MessageList []*common.MessageInfo `json:"message_list"`
}

// MessageActionParam 发送信息参数
type MessageActionParam struct {
	ToUserId int64  `json:"to_user_id" query:"to_user_id"`
	Action   int16  `json:"action_type" query:"action_type" vd:"@:$ >= 1 && $<=1; msg:'不支持的操作'"`
	Content  string `json:"content" query:"content" vd:"@:mblen($) >= 1 && mblen($) <= 256; msg:'消息长度在1-256之间'"`
}

// MessageActionResp 发送信息响应数据
type MessageActionResp struct {
	BaseResp
}
