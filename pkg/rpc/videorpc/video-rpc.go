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

// FavoriteVideo 调用点赞视频 RPC 服务
func FavoriteVideo(ctx context.Context, param *video.FavoriteVideoParam) error {
	resp, err := videoClient.FavoriteVideo(ctx, param)
	if err != nil {
		return err
	}
	return rpc.ParseRpcResponse(resp.BaseResp)
}

// FavoriteList 调用点赞列表 RPC 服务
func FavoriteList(ctx context.Context, param *video.FavoriteListParam) ([]*common.VideoInfo, error) {
	resp, err := videoClient.FavoriteList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.VideoList, nil
}

// CommentVideo 调用评论视频 RPC 服务
func CommentVideo(ctx context.Context, param *video.CommentVideoParam) (*common.CommentInfo, error) {
	resp, err := videoClient.CommentVideo(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.Comment, nil
}

// DeleteComment 调用删除评论 RPC 服务
func DeleteComment(ctx context.Context, param *video.DeleteCommentParam) error {
	resp, err := videoClient.DeleteComment(ctx, param)
	if err != nil {
		return err
	}
	return rpc.ParseRpcResponse(resp.BaseResp)
}

// CommentList 调用评论视频 RPC 服务
func CommentList(ctx context.Context, param *video.CommentListParam) ([]*common.CommentInfo, error) {
	resp, err := videoClient.CommentList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.CommentList, nil
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
