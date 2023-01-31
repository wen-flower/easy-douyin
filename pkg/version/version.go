package version

import (
	"fmt"
	"runtime"
)

// 使用 build -ldflag -X "imprtpath.name=value" 填充的值
var (
	Version      = "v0.0.0"              // Version 语义化版本号
	BuildDate    = "1970-01-01T00:00:00" // 构建时间, $(date -u +'%Y-%m-%dT%H:%M:%S') 命令的输出
	GitCommit    = ""                    // Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出
	GitTreeState = ""                    // 代表构建时 Git 仓库的状态，可能的值有：clean, dirty
)

// Info 包含了版本信息.
type Info struct {
	Version      string // 版本
	GitCommit    string // git rev-parse HEAD 的值
	GitTreeState string // 代表当前 Git 仓库的状态
	BuildDate    string // 构建时间
	GoVersion    string // Golang 的版本
	Compiler     string // 编译器
	Platform     string // 平台
}

func Get() Info {
	return Info{
		Version:      Version,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (info Info) String() string {
	return fmt.Sprintf(
		"Version      : %s\nGitCommit    : %s\nGitTreeState : %s\nBuildDate    : %s\nGoVersion    : %s\nCompiler     : %s\nPlatform     : %s",
		info.Version, info.GitCommit, info.GitTreeState, info.BuildDate, info.GoVersion, info.Compiler, info.Platform,
	)
}
