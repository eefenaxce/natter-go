# NatterGo

NatterGo 是一个基于 Go 语言开发的网络通信相关项目，支持多种网络协议和穿透能力，适合需要高性能、跨平台的网络应用开发者。

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

请根据 `cmd/` 目录下的子命令选择入口，常见用法：

```bash
go run ./cmd/xxx  # 替换xxx为具体子命令
```

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