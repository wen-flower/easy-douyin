package userrpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wen-flower/easy-douyin/cmd/user/consts"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/kitex_gen/user/userservice"
	"github.com/wen-flower/easy-douyin/pkg/constant"
	"github.com/wen-flower/easy-douyin/pkg/mw"
	"github.com/wen-flower/easy-douyin/pkg/rpc"
)

var userClient userservice.Client

// CreateUser 调用创建用户 RPC 服务
func CreateUser(ctx context.Context, param *user.CreateUserParam) (*int64, error) {
	resp, err := userClient.CreateUser(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserId, nil
}

// CheckUser 调用检查用户账号密码 RPC 服务
func CheckUser(ctx context.Context, param *user.CheckUserParam) (*int64, error) {
	resp, err := userClient.CheckUser(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserId, nil
}

// QueryUser 调用查询用户信息 RPC 服务
func QueryUser(ctx context.Context, param *user.QueryUserParam) ([]*common.UserInfo, error) {
	resp, err := userClient.QueryUser(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// FollowUser 调用关注用户 RPC 服务
func FollowUser(ctx context.Context, param *user.FollowUserParam) error {
	resp, err := userClient.FollowUser(ctx, param)
	if err != nil {
		return err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return err
	}
	return nil
}

// FollowList 调用查询用户关注列表 RPC 服务
func FollowList(ctx context.Context, param *user.FollowListParam) ([]*common.UserInfo, error) {
	resp, err := userClient.FollowList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// FollowerList 调用查询用户关注列表 RPC 服务
func FollowerList(ctx context.Context, param *user.FollowerListParam) ([]*common.UserInfo, error) {
	resp, err := userClient.FollowerList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// FriendList 调用查询用户好友列表 RPC 服务
func FriendList(ctx context.Context, param *user.FriendListParam) ([]*common.UserInfo, error) {
	resp, err := userClient.FriendList(ctx, param)
	if err != nil {
		return nil, err
	}
	if err = rpc.ParseRpcResponse(resp.BaseResp); err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// 初始化用户服务 PRC 客户端
func Init() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
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

	userClient = c
}
