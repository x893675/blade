# Blade

**Blade** 是一个基于 `kubebuilder` 扩展更改的脚手架工具，用来生成常用的 HTTP/GRPC 代码

> Blade 出自 `JRPG 异度之刃2` 中的 `异刃` 一词

## Installation

```shell
curl -sfL https://oss.hanamichi.wiki/install.sh | bash

mv blade /usr/local/bin/
```

## Getting Started

1. 初始化 HTTP 项目

   ```shell
   # 初始化项目，设置作者，仓库名，项目名
   blade init --owner "x893675" --repo github.com/x893675/blade-test --project-name blade-test
   ```
2. 添加 API

   ```shell
   # 添加 API, API 路由为 /foo/v1/user
   blade create api --group foo --version v1 --kind User
   # 在同一 GROUP 下添加另一个资源对象
   blade create api --group foo --version v1 --kind Group
   # 添加另一组 API
   blade create api --group bar --version v1 --kind User
   ```
3. `make help` 查看 Makefile 帮助信息

## Demo Video

[![Demo Video](https://asciinema.org/a/554608.svg)](https://asciinema.org/a/554608)

## Inspire By

- [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)
