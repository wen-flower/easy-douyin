package kitexlog

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/wen-flower/easy-douyin/pkg/mlog"
)

// Init 日志初始化
func Init(debug bool, text bool, pretty bool) {
	klog.SetLogger(kitexlogrus.NewLogger(
		kitexlogrus.WithLogger(mlog.Logger(
			mlog.ParseLevel(debug),
			mlog.ParseFormatter(text, pretty),
		)),
	))
	if debug {
		klog.SetLevel(klog.LevelDebug)
	} else {
		klog.SetLevel(klog.LevelInfo)
	}
}
