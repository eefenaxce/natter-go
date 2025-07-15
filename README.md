# NatterGo

NatterGo 是一个基于 Go 语言开发的轻量级端口转发与 NAT 穿透工具，支持 UDP/TCP，适用于内网穿透、端口映射等场景。

## 主要功能
- 支持 UDP/TCP 端口转发
- 支持 STUN 协议自动获取公网映射
- 支持 KeepAlive 保持连接活跃
- 支持多种转发方式（socket/iptables）
- 支持 UPnP 自动端口映射
- 丰富的日志输出，便于调试

## 安装

1. 克隆项目：
```bash
git clone https://github.com/eefenaxce/natter-go.git
cd natter-go
```
2. 构建：
```bash
go build -o nattergo ./cmd
```

## 使用方法

```bash
./nattergo [参数]
```
常用参数：
- `-i` 绑定IP（默认0.0.0.0）
- `-b` 绑定端口（0为自动）
- `-t` 目标IP
- `-p` 目标端口
- `-u` UDP模式（默认TCP）
- `-m` 转发方式（socket/iptables）
- `-k` keepalive间隔
- `-h` keepalive主机
- `-keep-port` keepalive端口
- `-U` 启用UPnP
- `-v` 输出详细日志

## 目录结构
```
cmd/            # 程序入口
internal/
  ├─config/     # 配置解析
  ├─forward/    # 转发实现
  ├─keepalive/  # KeepAlive 保活
  ├─logger/     # 日志模块
  └─stun/       # STUN 映射
```

## 依赖
- Go 1.24+
- github.com/pion/stun
- github.com/google/gopacket
- github.com/huin/goupnp

依赖详情见 `go.mod`。

## 许可证

本项目采用 [Apache License 2.0](LICENSE)。 
