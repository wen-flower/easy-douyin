# 简单的抖音后端服务

## 简介

使用 Hertz/Kitex/Gorm 实现一个基于微服务架构的简版抖音服务端

| 服务名          | 说明           | 框架          | 对外协议   | 路径        | IDL文件              |
|--------------|--------------|-------------|--------|-----------|--------------------|
| douyin-api   | 对外提供 HTTP 接口 | hertz/kitex | http   | cmd/api   | idl/api/*.thrift   |
| douyin-user  | 管理用户数据       | gorm/kitex  | thrift | cmd/user  | idl/user/*.thrift  |
| douyin-video | 管理视频数据       | gorm/kitex  | thrift | cmd/video | idl/video/*.thrift |

说明：
- 服务注册中心：[Etcd](https://etcd.io/)
- 服务与服务之间使用 `RPC` 的方式进行调用

## 基础特性

- Hertz：
    - 使用 Hertz 做服务聚合，对外提供 HTTP 接口
    - 使用 `obs-opentelemetry`、 `jarger for tracing`、 `metrics`、 `logging`
    - Middleware
        - JWT
        - pprof
        - requestid `TODO`
- Kitex：
    - 使用 Kitex 在服务内部交互
    - 使用 `obs-opentelemetry`、 `jarger for tracing`、 `metrics`、 `logging`
    - 使用 `registry-etcd` 做为服务发现和注册
- Gorm：
    - `使用 Gorm 进行数据库的 ORM 映射`
- Jaeger：
    - 链路追踪
- Grafana：
    - 数据仪表盘

## 目录介绍

[TODO](https://github.com/cloudwego/biz-demo/tree/main/easy_note#catalog-introduce)

## 快速开始

### 部署环境

```shell
docker-compose up
```

### 运行服务

API：
```shell
cd cmd/api
make build
_output/douyin-api --debug --port=380000 --log-pretty
```
User RPC：
```shell
cd cmd/user
make build
_output/douyin-user --debug --port=381000 --log-pretty
```
Video RPC：
```shell
cd cmd/video
make build
_output/douyin-video --debug --port=382000 --log-pretty
```
Chat RPC：
```shell
cd cmd/chat
make build
_output/douyin-chat --debug --port=383000 --log-pretty
```

### Jaeger
> 开源的、端到端的分布式链路追踪

在浏览器上访问 [http://127.0.0.1:16686/](http://127.0.0.1:16686/)

### Grafana
> 任何数据的仪表盘

在浏览器上访问 [http://127.0.0.1:3000/](http://127.0.0.1:3000/)

在 `configs` 文件夹中保存有简单配置的仪表盘 JSON 文件


