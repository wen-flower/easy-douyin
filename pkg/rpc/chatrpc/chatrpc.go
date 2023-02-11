package chatrpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wen-flower/easy-douyin/cmd/chat/consts"
	"github.com/wen-flower/easy-douyin/kitex_gen/chat"
	"github.com/wen-flower/easy-douyin/kitex_gen/chat/chatservice"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/pkg/mw"
	"github.com/wen-flower/easy-douyin/pkg/rpc"
)

var chatClient chatservice.Client

// MessageAction 调用发送消息 RPC 服务
func MessageAction(ctx context.Context, param *chat.MessageActionParam) error {
	resp, err := chatClient.MessageAction(ctx, param)
	if err != nil {
		return err
	}
	return rpc.ParseRpcResponse(resp.BaseResp)
}

// MessageList 调用获取聊天记录 RPC 服务
func MessageList(ctx context.Context, param *chat.MessageListParam) ([]*common.MessageInfo, error) {
	resp, err := chatClient.MessageList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.MessageList, nil
}

// Init 初始化用户服务 PRC 客户端
func Init(etcdAddress string) {
	r, err := etcd.NewEtcdResolver([]string{etcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := chatservice.NewClient(
		consts.ServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(rpc.ClientBasicInfo),
	)
	if err != nil {
		panic(err)
	}

	chatClient = c
}
