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
    - `TODO`
- Kitex：
    - `TODO`
- Gorm：
    - `TODO`
- Jaeger：
    - `TODO`
- Grafana：
    - `TODO`

## 目录介绍

[TODO](https://github.com/cloudwego/biz-demo/tree/main/easy_note#catalog-introduce)

## 快速开始

### 部署环境

`TODO`：使用 [Docker](https://github.com/docker/compose#docker-compose-v2)

### Jaeger
> 开源的、端到端的分布式链路追踪

在浏览器上访问 [http://127.0.0.1:16686/](http://127.0.0.1:16686/)

### Grafana
> 任何数据的仪表盘

在浏览器上访问 [http://127.0.0.1:3000/](http://127.0.0.1:3000/)


