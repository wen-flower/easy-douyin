#!/usr/bin/env zsh

# 检查参数是否传递了参数
if [ $# -ne 1 ]; then
  echo "使用代码生成脚本接受一个参数：$0 <idl文件名>"
  echo "eg: $0 user"
  exit 1
fi

# 将下划写或者中划线装驼峰形式的函数，接受一个参数
to_camel_case() {
  echo "$1" | sed -e "s/\b[a-z]/\u&/g" -e "s/[-_]//g"
}

# 将可能需要用到的变量全局化
# =======================
# 脚本的目录
SCRIPT_DIR=$(cd "$(dirname "$0")" || exit; pwd)
export SCRIPT_DIR

# 服务的包即命名空间，xxx
export SERVICE_PACKAGE="$1"

# 服务名，douyin-xxx
export SERVICE_NAME="douyin-$1"

# 服务接口名，XxxService
SERVICE_NAME_CAMEL_CASE=$(to_camel_case "$1-service")
export SERVICE_NAME_CAMEL_CASE

# 服务全小写名且没有其他符号，xxxservice
export SERVICE_NAME_LOWER="$1service"

# IDL 文件路径
export IDL_FILE="idl/$1-service.thrift"

# 对于服务的文件夹路径
CMD_DIR="$(pwd)/cmd/$1"
export CMD_DIR

# 检查 thrift idl 文件是否存在
if [ ! -f "$IDL_FILE" ]; then
  echo "文件 $IDL_FILE 不存在"
  exit 1
fi

kitex -thrift-plugin validator "$IDL_FILE" || exit
echo "客户端代码生成完成"

# 替换生成文件的所有变量，接受一个参数
replace_file_var() {
  sed -i \
    -e "s/#SERVICE_PACKAGE#/$SERVICE_PACKAGE/g" \
    -e "s/#SERVICE_NAME#/$SERVICE_NAME/g" \
    -e "s/#SERVICE_NAME_CAMEL_CASE#/$SERVICE_NAME_CAMEL_CASE/g" \
    -e "s/#SERVICE_NAME_LOWER#/$SERVICE_NAME_LOWER/g" \
    "$1"
#  sed -i "s/#SERVICE_PACKAGE#/$SERVICE_PACKAGE/g" "$1"
#  sed -i "s/#SERVICE_NAME#/$SERVICE_NAME/g" "$1"
#  sed -i "s/#SERVICE_NAME_CAMEL_CASE#/$SERVICE_NAME_CAMEL_CASE/g" "$1"
#  sed -i "s/#SERVICE_NAME_LOWER#/$SERVICE_NAME_LOWER/g" "$1"
}

# 遍历文件，然后替换，接受一个参数
replace_files_var() {
  cd "$CMD_DIR" || exit
  for file in **/* ; do
    if [ -f "$file" ]; then
      replace_file_var "$CMD_DIR/$file"
      if [[ "$file" == *"@go" ]]; then
        mv "$CMD_DIR/$file" "$CMD_DIR/${file//@go/.go}"
      fi
    fi
  done

  # 处理隐藏文件
  replace_file_var "$CMD_DIR/.air.toml"
}

# 创建服务相关的文件和目录
# ====================
# 检查需要生成代码的位置文件夹是否存在，不存在则创建
if [ ! -d "$CMD_DIR" ]; then
  echo "创建 $CMD_DIR" 文件夹
  mkdir -p "$CMD_DIR"
  # kitex -use github.com/wen-flower/easy-douyin/kitex_gen -service "douyin-$1" "../../$idlfile"
  cp -r "$SCRIPT_DIR/kitex/." "$CMD_DIR/"

  replace_files_var
  echo "微服务代码生成完成"
else
  echo "无法生成微服务代码: $CMD_DIR 文件夹存在"
fi
