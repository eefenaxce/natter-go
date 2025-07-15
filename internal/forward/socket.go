package forward

import (
	"fmt"
	"io"
	"net"
)

type SocketForwarder struct {
	listener net.Listener
}

func (f *SocketForwarder) Start(ip string, port int, toIP string, toPort int, udp bool) error {
	addr := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	var err error
	f.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := f.listener.Accept()
			if err != nil {
				return
			}
			go f.handle(conn, toIP, toPort)
		}
	}()
	return nil
}

func (f *SocketForwarder) handle(in net.Conn, toIP string, toPort int) {
	out, err := net.Dial("tcp", net.JoinHostPort(toIP, fmt.Sprintf("%d", toPort)))
	if err != nil {
		in.Close()
		return
	}
	go io.Copy(out, in)
	io.Copy(in, out)
}

func (f *SocketForwarder) Stop() {
	if f.listener != nil {
		f.listener.Close()
	}
}
