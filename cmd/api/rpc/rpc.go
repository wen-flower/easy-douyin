package rpc

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/pkg/constant"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

// Init 初始化 RPC 客户端
func Init() {
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	initUser()
}

var clientBasicInfo = &rpcinfo.EndpointBasicInfo{
	ServiceName: consts.ServiceName,
}

func parseRpcResponse(resp *common.BaseResp) error {
	if resp.Code != 0 {
		return errno.New(resp.Code, resp.Msg)
	}
	return nil
}
