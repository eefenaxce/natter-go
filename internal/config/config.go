package config

import (
	"flag"
	"time"
)

type Config struct {
	STUNList []string
	BindIP   string
	BindPort int
	ToIP     string
	ToPort   int
	UDP      bool
	Method   string
	Interval time.Duration
	KeepAddr string
	KeepPort int
	UPnP     bool
	Verbose  bool
}

func Parse() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.BindIP, "i", "0.0.0.0", "bind IP")
	flag.IntVar(&cfg.BindPort, "b", 0, "bind port (0=auto)")
	flag.StringVar(&cfg.ToIP, "t", "0.0.0.0", "target IP")
	flag.IntVar(&cfg.ToPort, "p", 0, "target port")
	flag.BoolVar(&cfg.UDP, "u", false, "UDP mode")
	flag.StringVar(&cfg.Method, "m", "socket", "forward method")
	flag.DurationVar(&cfg.Interval, "k", 15*time.Second, "keepalive interval")
	flag.StringVar(&cfg.KeepAddr, "h", "www.baidu.com", "keepalive host")
	flag.IntVar(&cfg.KeepPort, "keep-port", 80, "keepalive port")
	flag.BoolVar(&cfg.UPnP, "U", false, "enable UPnP")
	flag.BoolVar(&cfg.Verbose, "v", false, "verbose")
	flag.Parse()

	cfg.STUNList = []string{
		"stun.l.google.com:19302",
		"stun1.l.google.com:19302",
		"stun2.l.google.com:19302",
	}
	return cfg
}
