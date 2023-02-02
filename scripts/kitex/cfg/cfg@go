package cfg

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

// Port RPC 服务运行端口号
var Port int

// Debug 是否打开 Debug 级日志输出
var Debug bool

// LogJson 是否输出 JSON 格式
var LogJson bool

// LogPretty 是否输出彩色日志或者格式化 JSON
var LogPretty bool

// KitexRuntimeDir 	Kitex 的运行时位置
var KitexRuntimeDir string

// MySqlDNS MySQL 连接地址
var MySqlDNS string

func Init(flagSet *pflag.FlagSet) {
	flagSet.IntVarP(&Port, "port", "p", 8080, "指定服务运行的端口号")

	homeDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	flagSet.StringVar(&KitexRuntimeDir, "kitex-runtime-root", homeDir+"/_output", "指定 Kitex 运行目录")

	flagSet.StringVar(&MySqlDNS, "mysql-dns", "douyin:douyin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local", "MyQSL 连接地址")

	flagSet.BoolVar(&Debug, "debug", false, "是否打开 Debug 级日志输出")

	flagSet.BoolVar(&LogJson, "log-json", false, "是否输出 JSON 格式")

	flagSet.BoolVar(&LogPretty, "log-pretty", false, "是否输出彩色日志或者格式化 JSON")
}
