# NatterGo

NatterGo 是一个用 Go 语言实现的高性能网络通信工具。

## 项目简介

- 采用 Go 1.24 及以上版本
- 支持 UPnP、STUN、DTLS 等协议
- 依赖全部为宽松开源协议（MIT/BSD-2-Clause/BSD-3-Clause）

## 安装方法

```bash
git clone https://github.com/eefenaxce/natter-go.git
cd natter-go
go build ./cmd/...
```

## 使用方法

### 1. 编译所有命令行工具

```bash
go build -o bin ./cmd/...
```
编译后所有可执行文件会在 `bin/` 目录下。

### 2. 运行某个子命令

假设有一个子命令叫 `natter-server`，你可以这样运行：

```bash
./bin/natter-server --help
```

### 3. 常见命令与参数

以 `natter-server` 和 `natter-client` 为例（请根据你实际的cmd目录内容替换）：

#### 启动服务端
```bash
./bin/natter-server --listen :9000 --log-level debug
```
- `--listen`：监听地址和端口（如`:9000`）
- `--log-level`：日志级别（如`info`、`debug`）

#### 启动客户端
```bash
./bin/natter-client --server 127.0.0.1:9000 --user alice
```
- `--server`：服务端地址
- `--user`：用户名

### 4. 查看所有可用命令

```bash
ls ./cmd
```
每个子目录对应一个可执行程序，编译后在 `bin/` 目录下。

### 5. 查看命令帮助

```bash
./bin/xxx --help
```
会显示该命令支持的所有参数和用法。

### 6. 示例：UPnP端口映射

假如有 `natter-upnp` 工具：
```bash
./bin/natter-upnp --add --port 12345 --desc "NatterGo Test"
```

---

> 具体命令和参数请以 `cmd/` 目录下实际子命令和 `--help` 输出为准。

## 依赖说明

主要依赖如下：

- [github.com/huin/goupnp](https://github.com/huin/goupnp) (BSD-2-Clause)
- [github.com/pion/stun](https://github.com/pion/stun) (MIT)
- [github.com/pion/dtls/v2](https://github.com/pion/dtls) (MIT)
- [github.com/pion/logging](https://github.com/pion/logging) (MIT)
- [github.com/pion/transport/v2, v3](https://github.com/pion/transport) (MIT)
- [github.com/wlynxg/anet](https://github.com/wlynxg/anet) (MIT)
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) (BSD-3-Clause)
- [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) (BSD-3-Clause)
- [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) (BSD-3-Clause)
- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) (BSD-3-Clause)

所有依赖均为宽松许可证（MIT/BSD-2-Clause/BSD-3-Clause），可安全用于闭源或商业项目。

## 目录结构

```
NatterGo/
├── cmd/         # 各种命令行子程序入口
├── internal/    # 内部实现代码
├── go.mod       # Go module 文件
├── go.sum       # 依赖校验文件
```

## 许可证

本项目采用 BSD-2-Clause 许可证，详见下方：

```
Copyright (c) 2025, eefenaxce
All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

> 本项目所有依赖均为 MIT/BSD-2-Clause/BSD-3-Clause 等宽松许可证，无 GPL/LGPL 等传染性条款，可放心用于闭源或商业项目。 