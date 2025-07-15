package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eefenaxce/natter-go/internal/config"
	"github.com/eefenaxce/natter-go/internal/forward"
	"github.com/eefenaxce/natter-go/internal/keepalive"
	"github.com/eefenaxce/natter-go/internal/logger"
	"github.com/eefenaxce/natter-go/internal/stun"
)

func main() {
	// CLI 参数
	cfg := config.Parse()

	logger.Init(cfg.Verbose)

	// STUN 获取映射
	mapper := stun.New(cfg.STUNList, cfg.BindIP, cfg.BindPort, cfg.UDP)
	inner, outer, err := mapper.GetMapping()
	if err != nil {
		log.Fatalf("stun: %v", err)
	}

	logger.Infof("Mapped: %v -> %v", inner, outer)

	// 启动转发
	fw := forward.New(cfg.Method)
	if err := fw.Start(outer.IP, outer.Port, cfg.ToIP, cfg.ToPort, cfg.UDP); err != nil {
		log.Fatalf("forward: %v", err)
	}
	defer fw.Stop()

	// KeepAlive
	k := keepalive.New(cfg.KeepAddr, cfg.KeepPort, inner.IP, inner.Port, cfg.UDP)
	go k.Loop(cfg.Interval)

	// UPnP
	if cfg.UPnP {
		if err := forward.UPnPForward(outer.Port, inner.IP, inner.Port, cfg.UDP); err != nil {
			logger.Warnf("upnp: %v", err)
		}
	}

	// 优雅退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
