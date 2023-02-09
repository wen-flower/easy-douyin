package handler

import (
	"github.com/wen-flower/easy-douyin/cmd/chat/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/chat"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// 检查 ChatServiceImpl 是否实现了 chat.ChatService 接口
var _ chat.ChatService = (*ChatServiceImpl)(nil)

// 提取出错误的处理流程
func errProcess(baseResp **common.BaseResp, err *error) {
	var resp common.BaseResp
	if *err != nil {
		e := errno.ConvertErr(*err)
		resp.Msg = e.Msg()
		resp.Code = e.Code()
		*err = nil
	} else {
		resp.Code = errno.Success.Code()
		resp.Msg = errno.Success.Msg()
	}
	*baseResp = &resp
}

// 将 model.Message 列表转 common.MessageInfo 列表
func parseMessageInfoList(messages []*model.Message) []*common.MessageInfo {
	resp := make([]*common.MessageInfo, 0, len(messages))
	for _, message := range messages {
		resp = append(resp, &common.MessageInfo{
			Id:         message.Id,
			Content:    message.Content,
			CreateTime: message.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return resp
}
