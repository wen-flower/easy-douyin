package cfg

import "github.com/spf13/pflag"

// Port RPC 服务运行端口号
var Port int

// MaxRequestBodySize 请求 Body 最大允许的大小
var MaxRequestBodySize int

// Debug 是否打开 Debug 级日志输出
var Debug bool

// LogJson 是否输出 JSON 格式
var LogJson bool

// LogPretty 是否输出彩色日志或者格式化 JSON
var LogPretty bool

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

	flagSet.IntVar(&MaxRequestBodySize, "max-request-body", 100*1024*1204, "请求 Body 最大允许的大小(单位：byte)")

	flagSet.BoolVar(&Debug, "debug", false, "是否打开 Debug 级日志输出")

	flagSet.BoolVar(&LogJson, "log-json", false, "是否输出 JSON 格式")

	flagSet.BoolVar(&LogPretty, "log-pretty", false, "是否输出彩色日志或者格式化 JSON")
}
