# ==============================================================================
# 定义全局 Makefile 变量
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

# 版本包
VERSION_PACKAGE=github.com/wen-flower/easy-douyin/pkg/version
# 定义版本
VERSION := $(shell git describe --tags --always --match='v*')
# 仓库状态
GIT_TREE_STATE := "dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE = "clean"
endif
# 提交 SHA 值
GIT_COMMIT := $(shell git rev-parse HEAD)
# GO 构建 ldflags 参数
GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).Version=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%S')


# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: format build

# ==============================================================================
# 定义其他需要的伪目标

# 编译源码，依赖 tidy 目标自动添加/移除依赖包.
.PHONY: build
build:
	@go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/user-service $(ROOT_DIR)/main.go

# 格式化 Go 源码.
.PHONY: format
format:
	@gofmt -s -w ./

# 清理构建产物、临时文件等.
.PHONY: clean
clean:
	@-rm -vrf $(OUTPUT_DIR)