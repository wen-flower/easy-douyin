package gormlog

import (
	"github.com/sirupsen/logrus"
	"github.com/wen-flower/easy-douyin/pkg/mlog"
	"gorm.io/gorm/logger"
	gromlogrus "gorm.io/plugin/opentelemetry/logging/logrus"
	"time"
)

func GormLogger(json bool, pretty bool) logger.Interface {

	return logger.New(
		gromlogrus.NewWriter(gromlogrus.WithLogger(
			mlog.Logger(logrus.InfoLevel, mlog.ParseFormatter(json, pretty)),
		)),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      pretty,
			LogLevel:      logger.Info,
		},
	)
}
