package forward

import (
	"strings"
)

type Forwarder interface {
	Start(ip string, port int, toIP string, toPort int, udp bool) error
	Stop()
}

func New(method string) Forwarder {
	switch strings.ToLower(method) {
	case "socket":
		return &SocketForwarder{}
	case "iptables":
		return &IPTablesForwarder{}
	default:
		panic("unknown method: " + method)
	}
}
