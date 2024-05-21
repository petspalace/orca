package orca

import (
	"fmt"
	"log/slog"
	"net"
)

type Discovery struct {
	Territory Territory
}

func NewDiscovery(t Territory) *Discovery {
	d := Discovery{Territory: t}
	return &d
}

func (d *Discovery) Listen() error {
	slog.Info("start listening")

	adr, err := net.ResolveUDPAddr("udp4", ":7200")

	if err != nil {
		return err
	}

	con, err := net.ListenUDP("udp4", adr)

	if err != nil {
		return err
	}

	for {
		var buf [1024]byte

		_, radr, err := con.ReadFromUDP(buf[0:])

		if err != nil {
			return err
		}

		fmt.Printf("%v:%s\n", radr, string(buf[0:]))
	}

	return nil
}

func (d *Discovery) Broadcast() error {
	return nil
}
