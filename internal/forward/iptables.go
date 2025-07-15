package forward

import (
	"net"
	"os/exec"
	"strconv"
)

type IPTablesForwarder struct{}

func (f *IPTablesForwarder) Start(ip string, port int, toIP string, toPort int, udp bool) error {
	proto := "tcp"
	if udp {
		proto = "udp"
	}
	return exec.Command("iptables", "-t", "nat", "-A", "PREROUTING",
		"-p", proto, "-d", ip, "--dport", strconv.Itoa(port),
		"-j", "DNAT", "--to-destination", net.JoinHostPort(toIP, strconv.Itoa(toPort)),
	).Run()
}

func (f *IPTablesForwarder) Stop() {
	// TODO: 清理规则（需要记录规则句柄）
}
