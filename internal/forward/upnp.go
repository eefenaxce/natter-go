package forward

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

func UPnPForward(externalPort int, internalIP string, internalPort int, udp bool) error {
	ctx := context.Background()
	clients, _, errs := internetgateway1.NewWANIPConnection1Clients()
	if errs != nil && len(clients) == 0 {
		return fmt.Errorf("discover wan ip connection clients: %v", errs)
	}
	if len(clients) == 0 {
		return fmt.Errorf("no upnp internet gateway found")
	}
	c := clients[0]
	proto := "TCP"
	if udp {
		proto = "UDP"
	}
	ip := net.ParseIP(internalIP)
	if ip == nil {
		return fmt.Errorf("invalid internal IP: %s", internalIP)
	}
	err := c.AddPortMappingCtx(ctx, "", uint16(externalPort), proto, uint16(internalPort), ip.String(), true, "Natter", 0)
	if err != nil {
		return fmt.Errorf("add port mapping: %w", err)
	}
	log.Printf("[UPnP] Mapped %s:%d -> %s:%d", proto, externalPort, ip, internalPort)
	return nil
}
