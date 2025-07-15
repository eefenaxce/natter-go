package stun

import (
	"errors"
	"net"
	"time"

	"github.com/pion/stun"
)

type Addr struct {
	IP   string
	Port int
}

type Client struct {
	servers []string
	bindIP  string
	bindP   int
	udp     bool
}

func New(servers []string, bindIP string, bindP int, udp bool) *Client {
	return &Client{servers, bindIP, bindP, udp}
}

func (c *Client) GetMapping() (inner, outer Addr, err error) {
	conn, err := net.ListenPacket("udp4", net.JoinHostPort(c.bindIP, "0"))
	if err != nil {
		return
	}
	defer conn.Close()

	for _, srv := range c.servers {
		srvAddr, _ := net.ResolveUDPAddr("udp4", srv)
		// 1) 构造 BindingRequest
		msg := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
		// 2) 发送请求
		_, err := conn.WriteTo(msg.Raw, srvAddr)
		if err != nil {
			continue
		}
		// 3) 接收响应
		buf := make([]byte, 1500)
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, _, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}
		resp := &stun.Message{Raw: buf[:n]}
		if err := resp.Decode(); err != nil {
			continue
		}
		var xorAddr stun.XORMappedAddress
		if err = xorAddr.GetFrom(resp); err != nil {
			continue
		}
		laddr := conn.LocalAddr().(*net.UDPAddr)
		return Addr{laddr.IP.String(), laddr.Port},
			Addr{xorAddr.IP.String(), xorAddr.Port}, nil
	}
	return inner, outer, errors.New("no mapped address")
}
