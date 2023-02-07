package main

import (
	"fmt"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wen-flower/easy-douyin/cmd/video/cfg"
	"github.com/wen-flower/easy-douyin/cmd/video/consts"
	"github.com/wen-flower/easy-douyin/cmd/video/dal"
	"github.com/wen-flower/easy-douyin/cmd/video/handler"
	user "github.com/wen-flower/easy-douyin/kitex_gen/video/videoservice"
	"github.com/wen-flower/easy-douyin/pkg/command"
	"github.com/wen-flower/easy-douyin/pkg/constant"
	"github.com/wen-flower/easy-douyin/pkg/mlog/kitexlog"
	"github.com/wen-flower/easy-douyin/pkg/mw"
	"github.com/wen-flower/easy-douyin/pkg/rpc"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// 初始化数据层、日志框架等
func initialize() {
	dal.Init()

	rpc.Init(consts.ServiceName)
	userrpc.Init()

	kitexlog.Init(cfg.Debug, cfg.LogJson, cfg.LogPretty)
}

func run() error {
	// 初始化 otlp 跟踪和指标提供程序
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ServiceName),         // 配置服务名
		provider.WithExportEndpoint(constant.ExportEndpoint), // 配置导出地址
		provider.WithInsecure(),                              // 禁用导出程序 gRPC 的客户端传输安全性
	)

	// 创建 Etcd 注册中心信息
	r, err := etcd.NewEtcdRegistry([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// 解析用户服务运行的地址
	addr, err := net.ResolveTCPAddr(constant.TCP, fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}

	initialize()

	// 用户服务基础信息
	info := &rpcinfo.EndpointBasicInfo{
		ServiceName: consts.ServiceName,
	}

	svr := user.NewServer(new(handler.VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithMuxTransport(), // 将传输类型指定为多路复用。
		server.WithServerBasicInfo(info),
	)

	err = svr.Run()

	return err
}

// RPC 服务运行端口号
var port int

// Kitex 的运行时位置
var kitexRuntimeDir string

func main() {
	cmd := command.NewCommand(consts.ServiceName, func() error {
		if err := os.Setenv("KITEX_RUNTIME_ROOT", cfg.KitexRuntimeDir); err != nil {
			return err
		}
		if err := os.Setenv("KITEX_LOG_DIR", cfg.KitexRuntimeDir+"/log"); err != nil {
			return err
		}
		return run()
	})
	cfg.Init(cmd.PersistentFlags())
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
