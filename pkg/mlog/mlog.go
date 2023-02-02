package mlog

import (
	"github.com/sirupsen/logrus"
	"os"
)

var timestamp = "2006-01-02 15:04:05"

func Logger(level logrus.Level, formatter logrus.Formatter) *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    formatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}

func ParseLevel(debug bool) logrus.Level {
	if debug {
		return logrus.DebugLevel
	}
	return logrus.InfoLevel
}

func ParseFormatter(json bool, pretty bool) logrus.Formatter {
	if json {
		jsonFormatter := &logrus.JSONFormatter{
			TimestampFormat: timestamp,
			PrettyPrint:     pretty,
		}
		return jsonFormatter
	}
	textFormatter := &logrus.TextFormatter{
		ForceColors:     pretty,
		FullTimestamp:   true,
		TimestampFormat: timestamp,
	}
	return textFormatter
}
