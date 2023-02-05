package videorpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wen-flower/easy-douyin/cmd/video/consts"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/kitex_gen/video/videoservice"
	"github.com/wen-flower/easy-douyin/pkg/constant"
	"github.com/wen-flower/easy-douyin/pkg/mw"
	"github.com/wen-flower/easy-douyin/pkg/rpc"
)

var videoClient videoservice.Client

// CreateVideo 调用创建用 RPC 服务
func CreateVideo(ctx context.Context, param *video.CreateVideoParam) error {
	resp, err := videoClient.CreateVideo(ctx, param)
	if err != nil {
		return err
	}
	return rpc.ParseRpcResponse(resp.BaseResp)
}

// PublishList 调用发布列表 RPC 服务
func PublishList(ctx context.Context, param *video.PublishListParam) ([]*common.VideoInfo, error) {
	resp, err := videoClient.PublishList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.VideoList, nil
}

// VideoFeed 调用视频流 RPC 服务
func VideoFeed(ctx context.Context, param *video.VideoFeedParam) ([]*common.VideoInfo, error) {
	resp, err := videoClient.VideoFeed(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.VideoList, nil
}

// 初始化用户服务 PRC 客户端
func Init() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
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

	videoClient = c
}
