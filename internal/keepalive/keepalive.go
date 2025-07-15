package keepalive

import (
	"fmt"
	"net"
	"time"
)

type KeepAlive struct {
	dst  string
	conn net.Conn
	udp  bool
}

func New(host string, port int, srcIP string, srcPort int, udp bool) *KeepAlive {
	return &KeepAlive{
		dst: net.JoinHostPort(host, fmt.Sprintf("%d", port)),
	}
}

func (k *KeepAlive) Loop(interval time.Duration) {
	for {
		if k.udp {
			conn, _ := net.Dial("udp", k.dst)
			conn.Write([]byte{0})
			conn.Close()
		} else {
			conn, _ := net.Dial("tcp", k.dst)
			conn.Close()
		}
		time.Sleep(interval)
	}
}
