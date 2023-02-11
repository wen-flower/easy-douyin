package cfg

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
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

// RedisAddr Redis 连接地址
var RedisAddr string

// Etcd
var (
	// EtcdAddress Etcd 服务注册地址
	EtcdAddress string
)

// ExportEndpoint OTLP 暴露端口
var ExportEndpoint string

func Init(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&EtcdAddress, "etcd-address", "127.0.0.1:2379", "Etcd 服务地址")

	flagSet.StringVar(&ExportEndpoint, "export-endpoint", ":4317", "OTLP 暴露端点")

	flagSet.IntVarP(&Port, "port", "p", 8080, "指定服务运行的端口号")

	homeDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	flagSet.StringVar(&KitexRuntimeDir, "kitex-runtime-root", homeDir+"/_output", "指定 Kitex 运行目录")

	flagSet.StringVar(&MySqlDNS, "mysql-dns", "douyin:douyin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local", "MyQSL 连接地址")

	flagSet.StringVar(&RedisAddr, "redis-addr", "localhost:6379", "Redis 连接地址")

	flagSet.BoolVar(&Debug, "debug", false, "是否打开 Debug 级日志输出")

	flagSet.BoolVar(&LogJson, "log-json", false, "是否输出 JSON 格式")

	flagSet.BoolVar(&LogPretty, "log-pretty", false, "是否输出彩色日志或者格式化 JSON")
}
