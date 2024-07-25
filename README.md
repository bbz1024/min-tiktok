# min-tiktok

## 一. 项目简介

### 1.1 引言

> 本项目是基于青训营第六期项目min-tiktok实现，也是本人着手的第一个微服务项目，其中学习服务相关技术，Git Flow流，技术架构的设计。

#### 1.1.0 技术与框架

#### 1.1.1 接口文档与说明

[极简抖音App使用说明 - 第六届青训营版](https://bytedance.larkoffice.com/docx/NMneddpKCoXZJLxHePUcTzGgnmf)

[抖音项目方案说明-第六届青训营后端项目](https://bytedance.larkoffice.com/docx/BhEgdmoI3ozdBJxly71cd30vnRc)

## 二. 部署与运行

### 2.1 部署依赖服务

#### 方式一

本项目为了方便通过docker容器部署，所以需要先部署以下服务：

- MySQL
- Redis
- RabbitMQ
- Consul
- Gorse
- Nginx

#### 方式二

推荐使用docker-compose一键部署，无需手动部署服务，只需执行以下命令即可：

```bash
docker-compose up -d
```

### 2.2 修改配置文件

你需要修改每个服务下的etc/config.yaml文件，修改数据库连接信息，Redis连接信息，RabbitMQ连接信息，Consul连接信息，Gorse连接信息。

### 2.3 本地运行项目

为了方便在本项目下，你可以通过 [run.go](run.go) 运行项目，执行以下命令：

```go
go run run.go
```

> 注意：run.go会运行所有的服务。
--- 

### 2.4 运行项目

这里也提供了k8s部署方案，你可以通过 [k8s](k8s) 目录下的 yaml 文件来部署项目。
注意：在部署之前你需要上传镜像到你的镜像仓库，并修改 k8s 目录下的 yaml 文件中的镜像地址。

关于部署镜像你可以通过如下命令（[Dockerfile](Dockerfile)）：

```bash
docker build -t min-tiktok .

# push 操作
```

### 2.5 初始化数据

你需要运行

- [load_video_test](services/publish/load_video_test.go)
- [load_user_test](services/auths/load_user_test.go)




## 声明

本项目旨在分享学习内容，非盈利用途。所有内容均基于字节第六期青训营。如有侵权或引发争议，请及时联系作者，我们将尽快删除相关内容。
作者联系方式：2632141215@qq.com