#!/usr/bin/env zsh

if [ $# -ne 2 ]; then
  echo "参数错误, Usage: $0 <ServiceName> <TableName>"
  echo "ServiceName : 服务名，不要 douyin- 前缀"
  echo "TableName   : 表名，不需要 tb_ 前缀"
  exit 0
fi

# 将下划写或者中划线装驼峰形式的函数，接受一个参数
to_camel_case() {
  echo "$1" | sed -e "s/\b[a-z]/\u&/g" -e "s/[-_]//g"
}
SERVICE_NAME=$1
TABLE_NAME=$2
TABLE_NAME_CAMEL=$(to_camel_case "$TABLE_NAME")
DIR_PATH="cmd/$SERVICE_NAME/model"

if [ ! -d "$DIR_PATH" ]; then
    echo "服务 $SERVICE_NAME 文件夹 $DIR_PATH 不存在"
    exit 1
fi

cd "$DIR_PATH" || exit

db2struct --host=127.0.0.1 --mysql_port=3306 --user=douyin --password=douyin \
 --database=douyin --table="tb_$TABLE_NAME" \
 --gorm --no-json \
 --package=model --struct "$TABLE_NAME_CAMEL" \
 --target="$TABLE_NAME.go"

 sed -i -e "s/sets the insert table name for this struct type/结构体对应的数据库表名/" "$TABLE_NAME.go"

# 判断是否需要添加 time 包
 if grep -q "time.Time" "$TABLE_NAME.go"; then
   sed -i "2a import \"time\"\n" "$TABLE_NAME.go"
 fi

# 生成类名对应的常量
COLUMN_LIST=$(grep -E "^(\s+\S+){2}\s+\`\S+\`" "$TABLE_NAME.go" | \
  sed -E "s/^\s+(\S+)\s+\S+\s+\`.*column:([a-Z_]+).*\`.*$/\t$TABLE_NAME_CAMEL\1 = \"\2\"/g")

{
  echo ""
  echo "// 结构体字段对应的数据表列名常量"
  echo "const ("
  echo "$COLUMN_LIST"
  echo ")"
} >> "$TABLE_NAME.go"

gofmt -s -w "./$TABLE_NAME.go"