package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/wen-flower/easy-douyin/cmd/api/cfg"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/cmd/api/mw"
	"github.com/wen-flower/easy-douyin/cmd/api/router"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/pkg/command"
	"github.com/wen-flower/easy-douyin/pkg/cos"
	"github.com/wen-flower/easy-douyin/pkg/mlog/hertzlog"
	"github.com/wen-flower/easy-douyin/pkg/mlog/kitexlog"
	"github.com/wen-flower/easy-douyin/pkg/rpc"
	"github.com/wen-flower/easy-douyin/pkg/rpc/chatrpc"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// 初始化 RPC 客户端、日志框架等
func initialize() {
	rpc.Init(consts.ServiceName)
	userrpc.Init()
	videorpc.Init()
	chatrpc.Init()

	mw.InitJWT()

	cos.Init()

	utils.InitHertzValidator()

	hertzlog.Init(cfg.Debug, cfg.LogJson, cfg.LogPretty)
	kitexlog.Init(cfg.Debug, cfg.LogJson, cfg.LogPretty)
}

// 关闭数据库连接等
func shutdown() {
}

func run() error {
	initialize()

	tracer, config := tracing.NewServerTracer()

	h := server.New(
		server.WithMaxRequestBodySize(cfg.MaxRequestBodySize),
		server.WithHostPorts(fmt.Sprintf(":%d", cfg.Port)),
		tracer,
	)

	// 使用 pprof
	pprof.Register(h)

	// 使用 otel 中间件
	h.Use(tracing.ServerMiddleware(config))

	// 注册路由
	router.Register(h)

	// 启动，并添加停止时操作
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		shutdown()
	})
	h.Spin()

	return nil
}

func main() {
	cmd := command.NewCommand(consts.ServiceName, func() error {
		return run()
	})

	cfg.Init(cmd.PersistentFlags())

	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
