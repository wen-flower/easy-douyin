package hertzlog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/wen-flower/easy-douyin/pkg/mlog"
)

// Init 日志初始化
func Init(debug bool, json bool, pretty bool) {
	hlog.SetLogger(hertzlogrus.NewLogger(
		hertzlogrus.WithLogger(mlog.Logger(
			mlog.ParseLevel(debug),
			mlog.ParseFormatter(json, pretty),
		)),
	))
	if debug {
		hlog.SetLevel(hlog.LevelDebug)
	} else {
		hlog.SetLevel(hlog.LevelInfo)
	}
}
